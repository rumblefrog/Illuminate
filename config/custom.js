/**
 * Custom configuration
 * (sails.config.custom)
 *
 * One-off settings specific to your application.
 *
 * For more information on custom configuration, visit:
 * https://sailsjs.com/config/custom
 */

module.exports.custom = {

  /***************************************************************************
  *                                                                          *
  * Any other custom config this Sails app should use during development.    *
  *                                                                          *
  ***************************************************************************/
  // mailgunDomain: 'transactional-mail.example.com',
  // mailgunSecret: 'key-testkeyb183848139913858e8abd9a3',
  // stripeSecret: 'sk_test_Zzd814nldl91104qor5911gjald',
  // …

  maxUploadBytes: 20000000, // 20 MB
  
  minioEndpoint: 'play.minio.io',
  minioPort: 9000,
  minioSecure: true,
  minioAccessKey: 'Q3AM3UQ867SPQQA43P2F',
  minioSecretKey: 'zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG',
  minioBucket: 'Illuminate',


};
