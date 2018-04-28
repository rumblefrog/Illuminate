const mime = require('mime-types')
    , path = require('path')
    , { to } = require('await-to-js');

module.exports = {


  friendlyName: 'View',


  description: 'View something.',


  inputs: {

    id: {
      friendlyName: 'ID',
      type: 'string',
      required: true
    }

  },


  exits: {

    success: {
      responseType: "",
      statusCode: 200
    },

    invalidID: {
      responseType: "",
      statusCode: 400
    },

    notFound: {
      responseType: "",
      statusCode: 404
    },

    fetchFail: {
      responseType: "",
      statusCode: 500
    },

    updateFail: {
      responseType: "",
      statusCode: 500
    }

  },


  fn: async function (inputs, exits) {

    let err, result, stream, 
    needle = path.parse(inputs.id).name,
    datastore = sails.getDatastore();

    if (!datastore.driver.mongodb.ObjectID.isValid(needle))
      return exits.invalidID('Invalid ID');

    [ err, result ] = await to(
      datastore.manager
        .collection(sails.config.custom.collection)
        .findOne({
          _id: datastore.driver.mongodb.ObjectID(needle)
        })
    );

    if (err)
      return exits.fetchFail('Unable to fetch from database');

    if (result === null)
      return exits.notFound('ID not found');

    [ err ] = await to(
      datastore.manager
        .collection(sails.config.custom.collection)
        .updateOne(
          { _id: datastore.driver.mongodb.ObjectID(needle) },
          { $inc: { views: 1 } }
        )
    );

    //TODO: Append to logs:[] array

    if (err)
      return exits.updateFail('Failed to update document');

    [ err, stream ] = await to(
      sails.config.minio.Minio
        .getObject(
          sails.config.minio.settings.bucket,
          result.object
        )
    );

    if (err)
      return exits.fetchFail('Unable to fetch from storage');

    stream.pipe(this.res);
  }


};
