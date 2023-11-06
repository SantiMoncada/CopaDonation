const path = require("path");

module.exports = {
  mode: "production",
  entry: "./js/index.js",
  output: {
    path: path.resolve(__dirname, "public/dist"),
    filename: "index.js",
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
