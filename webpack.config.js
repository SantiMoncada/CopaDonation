const path = require("path");

module.exports = {
  entry: "./js/index.js",
  output: {
    path: path.resolve(__dirname, "public/scripts"),
    filename: "index.js",
  },
  mode: "production",
};
