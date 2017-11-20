const path = require('path');

module.exports = {
  entry: './browser/src/index.js',
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'browser/dist')
  }
};
