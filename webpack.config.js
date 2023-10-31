const path = require('path');

module.exports = {
    entry: './JS/index.js',
    output: {
        path: path.resolve(__dirname, 'assets/scripts'),
        filename: 'index.js',
    },
    mode: 'production',
};