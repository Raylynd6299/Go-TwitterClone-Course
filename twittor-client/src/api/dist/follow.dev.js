"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.checkFollow = checkFollow;
exports.followUserApi = followUserApi;
exports.unfollowUserApi = unfollowUserApi;

var _constant = require("../utils/constant");

var _auth = require("./auth");

function checkFollow(idUser) {
  var url = "".concat(_constant.API_HOST, "/consultaRelacion?id=").concat(idUser);
  var params = {
    headers: {
      Authorization: "Bearer ".concat((0, _auth.getTokenApi)())
    }
  };
  return fetch(url, params).then(function (response) {
    return response.json();
  }).then(function (result) {
    return result;
  })["catch"](function (err) {
    return err;
  });
}

function followUserApi(idUser) {
  var url = "".concat(_constant.API_HOST, "/altaRelacion?id=").concat(idUser);
  var params = {
    method: "POST",
    headers: {
      Authorization: "Bearer ".concat((0, _auth.getTokenApi)())
    }
  };
  return fetch(url, params).then(function (response) {
    return response.json();
  }).then(function (result) {
    return result;
  })["catch"](function (err) {
    return err;
  });
}

function unfollowUserApi(idUser) {
  var url = "".concat(_constant.API_HOST, "/bajaRelacion?id=").concat(idUser);
  var params = {
    method: "DELETE",
    headers: {
      Authorization: "Bearer ".concat((0, _auth.getTokenApi)())
    }
  };
  return fetch(url, params).then(function (response) {
    return response.json();
  }).then(function (result) {
    return result;
  })["catch"](function (err) {
    return err;
  });
}