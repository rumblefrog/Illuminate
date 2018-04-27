const mime = require('mime-types')
    , path = require('path')
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
    }

  },


  fn: async function (inputs, exits) {

    inputs.payload.upload({

      dirname: 'assets/images/',

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

      let err, exists, etag;

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
          mime.lookup(uploadedFiles[0].fd)
        )
      );

      if (err) {
        sails.log.error('Failed to put object')
        return exits.uploadFail('Failed to put object');
      }

      sails.log.info('Successfully put object');

      return exits.success();
    });
  }
};
