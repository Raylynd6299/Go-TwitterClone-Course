"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.getUserApi = getUserApi;

var _constant = require("../utils/constant");

var _auth = require("./auth");

function getUserApi(id) {
  var url = "".concat(_constant.API_HOST, "/verperfil?id=").concat(id);
  var params = {
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer ".concat((0, _auth.getTokenApi)())
    }
  };
  return fetch(url, params).then(function (response) {
    if (response.status >= 400) throw null;
    return response.json();
  }).then(function (result) {
    return result;
  })["catch"](function (err) {
    return err;
  });
}