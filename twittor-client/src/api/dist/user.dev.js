"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.getUserApi = getUserApi;
exports.uploadBannerApi = uploadBannerApi;
exports.uploadAvatarApi = uploadAvatarApi;
exports.updateInfoApi = updateInfoApi;

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
    if (response.status >= 400) {
      return null;
    }

    ;
    return response.json();
  }).then(function (result) {
    return result;
  })["catch"](function (err) {
    return err;
  });
}

function uploadBannerApi(file) {
  var url = "".concat(_constant.API_HOST, "/subirBanner");
  var formData = new FormData();
  formData.append("banner", file);
  var params = {
    method: "POST",
    headers: {
      Authorization: "Bearer ".concat((0, _auth.getTokenApi)())
    },
    body: formData
  };
  return fetch(url, params).then(function (response) {
    return response;
  }).then(function (result) {
    return result;
  })["catch"](function (err) {
    return err;
  });
}

function uploadAvatarApi(file) {
  var url = "".concat(_constant.API_HOST, "/subirAvatar");
  var formData = new FormData();
  formData.append("avatar", file);
  var params = {
    method: "POST",
    headers: {
      Authorization: "Bearer ".concat((0, _auth.getTokenApi)())
    },
    body: formData
  };
  return fetch(url, params).then(function (response) {
    return response;
  }).then(function (result) {
    return result;
  })["catch"](function (err) {
    return err;
  });
}

function updateInfoApi(data) {
  var url = "".concat(_constant.API_HOST, "/modificarPerfil");
  var params = {
    method: "PUT",
    headers: {
      Authorization: "Bearer ".concat((0, _auth.getTokenApi)())
    },
    body: JSON.stringify(data)
  };
  return fetch(url, params).then(function (response) {
    return response;
  })["catch"](function (err) {
    return err;
  });
}