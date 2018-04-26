module.exports = {


  friendlyName: 'Create',


  description: 'Create upload.',

  files: ['payload'],


  inputs: {

    key: {
      friendlyName: 'API Key',
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

    }, (err, uploadedFiles) => {

      if (err)
        return exits.uploadFail('Unable to save to asset directory');

      Minio.bucketExists(sails.config.custom.minioBucket, (err, exists) => {
        if (err)
          return exits.uploadFail('Unable to fetch bucket');

        if (!exists) {
          Minio.makeBucket(sails.config.custom.minioBucket, (err) => {
            if (err)
              return exits.uploadFail('Unable to create bucket');

            
          })
        }
      })
    
    });

    return exits.success();

  }


};
