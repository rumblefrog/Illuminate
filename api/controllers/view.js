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

    notFound: {
      responseType: "",
      statusCode: 404
    }

  },


  fn: async function (inputs, exits) {

    let err, result, 
    needle = path.parse(inputs.id).name,
    datastore = sails.getDatastore();

    [ err, result ] = await to(
      datastore
        .manager
        .collection(sails.config.custom.collection)
        .findOne({
          _id: datastore.driver.mongodb.ObjectID(needle)
        })
    );

    console.log(needle);
    console.log(err, result);

    return exits.success();

  }


};
