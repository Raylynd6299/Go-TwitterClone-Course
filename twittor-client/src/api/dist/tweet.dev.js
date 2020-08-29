"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.addTweetApi = addTweetApi;
exports.getUserTweetsApi = getUserTweetsApi;
exports.getTweetsFollowersApi = getTweetsFollowersApi;

var _constant = require("../utils/constant");

var _auth = require("./auth");

function addTweetApi(message) {
  var url = "".concat(_constant.API_HOST, "/tweet");
  var data = {
    mensaje: message
  };
  var params = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer ".concat((0, _auth.getTokenApi)())
    },
    body: JSON.stringify(data)
  };
  return fetch(url, params).then(function (response) {
    if (response.status >= 200 && response.status < 300) {
      return {
        code: response.status,
        message: "Tweet enviado."
      };
    }

    return {
      code: 500,
      message: "Error del servidor"
    };
  })["catch"](function (err) {
    return err;
  });
}

function getUserTweetsApi(idUser, page) {
  var url = "".concat(_constant.API_HOST, "/leoTweets?id=").concat(idUser, "&pagina=").concat(page);
  var params = {
    headers: {
      "Content-Type": "application/json",
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

function getTweetsFollowersApi() {
  var page = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : 1;
  var url = "".concat(_constant.API_HOST, "/leoTweetsSeguidores?pagina=").concat(page);
  var params = {
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer ".concat((0, _auth.getTokenApi)())
    }
  };
  return fetch(url, params).then(function (response) {
    return response.json();
  })["catch"](function (err) {
    return err;
  });
}