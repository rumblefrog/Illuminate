module.exports = async function (req, res, proceed) {

    if (req.param('key') !== sails.config.custom.key)
        return res.forbidden();

    return proceed();
}