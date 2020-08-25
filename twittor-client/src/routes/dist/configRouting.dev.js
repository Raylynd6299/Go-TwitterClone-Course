"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _Home = _interopRequireDefault(require("../page/Home"));

var _User = _interopRequireDefault(require("../page/User"));

var _Error = _interopRequireDefault(require("../page/Error404"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

var _default = [{
  path: "/:id",
  exact: true,
  page: _User["default"]
}, {
  path: "/",
  exact: true,
  page: _Home["default"]
}, {
  path: "*",
  page: _Error["default"]
}];
exports["default"] = _default;