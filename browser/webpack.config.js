const path = require('path');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');

module.exports = {
  entry: {
    index: './src/components/app/index.js'
  },
  output: {
    filename: 'static/[name].bundle.js',
    path: path.resolve(__dirname, 'www')
  },
  devtool: 'inline-source-map',
  plugins: [
    new CleanWebpackPlugin(['www/static']),
    new CopyWebpackPlugin([
      { from: path.resolve(__dirname, 'node_modules/@webcomponents/webcomponentsjs/webcomponents-hi-sd-ce.js'), to: 'static/webcomponents-hi-sd-ce.js' },
      { from: path.resolve(__dirname, 'node_modules/web-animations-js/web-animations-next-lite.min.js'), to: 'static/web-animations-next-lite.min.js' },
      // style.css is built from the VSCode task 'build sass' as webpack is only for ES modules
      { from: path.resolve(__dirname, 'dist/style.css'), to: 'static/style.css' },      
    ])
  ],
  module: {
    rules: [
      {
        test: /\.html$/,
        use: ['text-loader']
      },
      {
        test: /\.scss$/,
        use: [ 'css-loader', 'sass-loader']  
      }
    ]
  }
};

