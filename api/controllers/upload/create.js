const mime = require('mime-types')
    , path = require('path')
    , fs = require('fs')
    , { to } = require('await-to-js');

module.exports = {


  friendlyName: 'Create',


  description: 'Create upload.',

  files: ['payload'],


  inputs: {

    key: {
      friendlyName: 'API Key',
      type: 'string',
      required: true
    },

    payload: {
      example: '===',
      required: true
    }


  },


  exits: {

    success: {
      responseType: "",
      statusCode: 201
    },

    invalidPayload: {
      responseType: "",
      statusCode: 400
    },

    uploadFail: {
      responseType: "",
      statusCode: 500
    },

    insertionFail: {
      responseType: "",
      statusCode: 500
    }

  },


  fn: async function (inputs, exits) {

    inputs.payload.upload({

      maxBytes: sails.config.custom.maxUploadBytes

    }, async (e, uploadedFiles) => {

      if (e) {
        sails.log.error('Unable to save to assets directory');
        return exits.uploadFail('Unable to save to asset directory');
      }

      if (uploadedFiles.length === 0) {
        sails.log.error('No file uploaded');
        return exits.uploadFail('No file uploaded');
      }

      let err, exists, etag, result, contentType = mime.lookup(uploadedFiles[0].fd);

      [ err, exists ] = await to(
        sails.config.minio.Minio.bucketExists(
          sails.config.minio.settings.bucket
        )
      );

      if (err) {
        sails.log.error('Unable to fetch bucket');
        return exits.uploadFail('Unable to fetch bucket');
      }

      if (!exists) {
        [ err ] = await to(
          sails.config.minio.Minio.makeBucket(
            sails.config.minio.settings.bucket
          )
        );

        if (err) {
          sails.log.error('Unable to create bucket');
          return exits.uploadFail('Unable to create bucket');
        }

        sails.log.info('Created bucket');
      }

      [ err, etag ] = await to(
        sails.config.minio.Minio.fPutObject(
          sails.config.minio.settings.bucket,
          path.basename(uploadedFiles[0].fd),
          uploadedFiles[0].fd,
          contentType
        )
      );

      if (err) {
        sails.log.error('Failed to put object')
        return exits.uploadFail('Failed to put object');
      }

      [ err, result ] = await to(
        sails.getDatastore()
          .manager
          .collection(sails.config.custom.collection)
          .insert({
            etag: etag,
            object: path.basename(uploadedFiles[0].fd),
            contentType: contentType,
            views: 0,
            logs: []
          })
      );

      if (err) {
        sails.log.error('Unable to insert into MongoDB');
        return exits.insertionFail('Unable to insert into MongoDB');
      }

      fs.unlinkSync(uploadedFiles[0].fd);

      sails.log.info('Successfully put object', etag);

      return exits.success(`${result.insertedIds[0]}.${mime.extension(contentType)}`);
    });
  }
};
