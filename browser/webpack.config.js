const path = require('path');
const CleanWebpackPlugin = require('clean-webpack-plugin');

module.exports = {
  entry: {
    index: './src/index.jsx',
  },
  devtool: 'inline-source-map',
  plugins: [
    new CleanWebpackPlugin(['www'])
  ],
  output: {
    filename: 'static/[name].bundle.js',
    path: path.resolve(__dirname, 'www')
  },
  module: {
    loaders: [
      {
        test: /.jsx?$/,
        loader: 'babel-loader',
        exclude: /node_modules/,
        query: {
          presets: ['es2015', 'react']
        }
      }
    ]
  },
  resolve: {
    extensions: ['.js', '.jsx']
  }
};