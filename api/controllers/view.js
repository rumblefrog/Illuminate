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

    

    return exits.success();

  }


};
