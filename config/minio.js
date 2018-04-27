const SDK = require('minio');

module.exports.minio = {

    settings: {

      bucket: 'illuminate' // Must be lowercase or else it will error

    },

    Minio: new SDK.Client({

      endPoint: 'play.minio.io',
      port: 9000,
      secure: true,
      accessKey: 'Q3AM3UQ867SPQQA43P2F', 
      secretKey: 'zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG',

  }),
}