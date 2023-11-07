const path = require("path");

module.exports = {
  mode: "production",
  entry: {
    main: './js/main.js',
    startup: './js/startup.js'
  },
  output: {
    path: path.resolve(__dirname, "public/dist"),
    filename: '[name].js',
  },
  module: {
    rules: [
      {
        test: /\.css$/i,
        include: path.resolve(__dirname, "css"),
        use: ["style-loader", "css-loader", "postcss-loader"],
      },
    ],
  },
};
