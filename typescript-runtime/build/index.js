function _arrayLikeToArray(r, a) {
  (null == a || a > r.length) && (a = r.length);
  for (var e = 0, n = Array(a); e < a; e++) n[e] = r[e];
  return n;
}
function _arrayWithHoles(r) {
  if (Array.isArray(r)) return r;
}
function _arrayWithoutHoles(r) {
  if (Array.isArray(r)) return _arrayLikeToArray(r);
}
function _assertClassBrand(e, t, n) {
  if ("function" == typeof e ? e === t : e.has(t)) return arguments.length < 3 ? t : n;
  throw new TypeError("Private element is not present on this object");
}
function _assertThisInitialized(e) {
  if (void 0 === e) throw new ReferenceError("this hasn't been initialised - super() hasn't been called");
  return e;
}
function _callSuper(t, o, e) {
  return o = _getPrototypeOf(o), _possibleConstructorReturn(t, _isNativeReflectConstruct() ? Reflect.construct(o, e || [], _getPrototypeOf(t).constructor) : o.apply(t, e));
}
function _checkPrivateRedeclaration(e, t) {
  if (t.has(e)) throw new TypeError("Cannot initialize the same private elements twice on an object");
}
function _classCallCheck(a, n) {
  if (!(a instanceof n)) throw new TypeError("Cannot call a class as a function");
}
function _classPrivateFieldGet2(s, a) {
  return s.get(_assertClassBrand(s, a));
}
function _classPrivateFieldInitSpec(e, t, a) {
  _checkPrivateRedeclaration(e, t), t.set(e, a);
}
function _classPrivateFieldSet2(s, a, r) {
  return s.set(_assertClassBrand(s, a), r), r;
}
function _construct(t, e, r) {
  if (_isNativeReflectConstruct()) return Reflect.construct.apply(null, arguments);
  var o = [null];
  o.push.apply(o, e);
  var p = new (t.bind.apply(t, o))();
  return r && _setPrototypeOf(p, r.prototype), p;
}
function _defineProperties(e, r) {
  for (var t = 0; t < r.length; t++) {
    var o = r[t];
    o.enumerable = o.enumerable || !1, o.configurable = !0, "value" in o && (o.writable = !0), Object.defineProperty(e, _toPropertyKey(o.key), o);
  }
}
function _createClass(e, r, t) {
  return r && _defineProperties(e.prototype, r), t && _defineProperties(e, t), Object.defineProperty(e, "prototype", {
    writable: !1
  }), e;
}
function _defineProperty(e, r, t) {
  return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, {
    value: t,
    enumerable: !0,
    configurable: !0,
    writable: !0
  }) : e[r] = t, e;
}
function _getPrototypeOf(t) {
  return _getPrototypeOf = Object.setPrototypeOf ? Object.getPrototypeOf.bind() : function (t) {
    return t.__proto__ || Object.getPrototypeOf(t);
  }, _getPrototypeOf(t);
}
function _inherits(t, e) {
  if ("function" != typeof e && null !== e) throw new TypeError("Super expression must either be null or a function");
  t.prototype = Object.create(e && e.prototype, {
    constructor: {
      value: t,
      writable: !0,
      configurable: !0
    }
  }), Object.defineProperty(t, "prototype", {
    writable: !1
  }), e && _setPrototypeOf(t, e);
}
function _isNativeFunction(t) {
  try {
    return -1 !== Function.toString.call(t).indexOf("[native code]");
  } catch (n) {
    return "function" == typeof t;
  }
}
function _isNativeReflectConstruct() {
  try {
    var t = !Boolean.prototype.valueOf.call(Reflect.construct(Boolean, [], function () {}));
  } catch (t) {}
  return (_isNativeReflectConstruct = function () {
    return !!t;
  })();
}
function _iterableToArray(r) {
  if ("undefined" != typeof Symbol && null != r[Symbol.iterator] || null != r["@@iterator"]) return Array.from(r);
}
function _iterableToArrayLimit(r, l) {
  var t = null == r ? null : "undefined" != typeof Symbol && r[Symbol.iterator] || r["@@iterator"];
  if (null != t) {
    var e,
      n,
      i,
      u,
      a = [],
      f = !0,
      o = !1;
    try {
      if (i = (t = t.call(r)).next, 0 === l) ; else for (; !(f = (e = i.call(t)).done) && (a.push(e.value), a.length !== l); f = !0);
    } catch (r) {
      o = !0, n = r;
    } finally {
      try {
        if (!f && null != t.return && (u = t.return(), Object(u) !== u)) return;
      } finally {
        if (o) throw n;
      }
    }
    return a;
  }
}
function _nonIterableRest() {
  throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.");
}
function _nonIterableSpread() {
  throw new TypeError("Invalid attempt to spread non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.");
}
function ownKeys(e, r) {
  var t = Object.keys(e);
  if (Object.getOwnPropertySymbols) {
    var o = Object.getOwnPropertySymbols(e);
    r && (o = o.filter(function (r) {
      return Object.getOwnPropertyDescriptor(e, r).enumerable;
    })), t.push.apply(t, o);
  }
  return t;
}
function _objectSpread2(e) {
  for (var r = 1; r < arguments.length; r++) {
    var t = null != arguments[r] ? arguments[r] : {};
    r % 2 ? ownKeys(Object(t), !0).forEach(function (r) {
      _defineProperty(e, r, t[r]);
    }) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : ownKeys(Object(t)).forEach(function (r) {
      Object.defineProperty(e, r, Object.getOwnPropertyDescriptor(t, r));
    });
  }
  return e;
}
function _possibleConstructorReturn(t, e) {
  if (e && ("object" == typeof e || "function" == typeof e)) return e;
  if (void 0 !== e) throw new TypeError("Derived constructors may only return object or undefined");
  return _assertThisInitialized(t);
}
function _setPrototypeOf(t, e) {
  return _setPrototypeOf = Object.setPrototypeOf ? Object.setPrototypeOf.bind() : function (t, e) {
    return t.__proto__ = e, t;
  }, _setPrototypeOf(t, e);
}
function _slicedToArray(r, e) {
  return _arrayWithHoles(r) || _iterableToArrayLimit(r, e) || _unsupportedIterableToArray(r, e) || _nonIterableRest();
}
function _toConsumableArray(r) {
  return _arrayWithoutHoles(r) || _iterableToArray(r) || _unsupportedIterableToArray(r) || _nonIterableSpread();
}
function _toPrimitive(t, r) {
  if ("object" != typeof t || !t) return t;
  var e = t[Symbol.toPrimitive];
  if (void 0 !== e) {
    var i = e.call(t, r);
    if ("object" != typeof i) return i;
    throw new TypeError("@@toPrimitive must return a primitive value.");
  }
  return (String )(t);
}
function _toPropertyKey(t) {
  var i = _toPrimitive(t, "string");
  return "symbol" == typeof i ? i : i + "";
}
function _typeof(o) {
  "@babel/helpers - typeof";

  return _typeof = "function" == typeof Symbol && "symbol" == typeof Symbol.iterator ? function (o) {
    return typeof o;
  } : function (o) {
    return o && "function" == typeof Symbol && o.constructor === Symbol && o !== Symbol.prototype ? "symbol" : typeof o;
  }, _typeof(o);
}
function _unsupportedIterableToArray(r, a) {
  if (r) {
    if ("string" == typeof r) return _arrayLikeToArray(r, a);
    var t = {}.toString.call(r).slice(8, -1);
    return "Object" === t && r.constructor && (t = r.constructor.name), "Map" === t || "Set" === t ? Array.from(r) : "Arguments" === t || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t) ? _arrayLikeToArray(r, a) : void 0;
  }
}
function _wrapNativeSuper(t) {
  var r = "function" == typeof Map ? new Map() : void 0;
  return _wrapNativeSuper = function (t) {
    if (null === t || !_isNativeFunction(t)) return t;
    if ("function" != typeof t) throw new TypeError("Super expression must either be null or a function");
    if (void 0 !== r) {
      if (r.has(t)) return r.get(t);
      r.set(t, Wrapper);
    }
    function Wrapper() {
      return _construct(t, arguments, _getPrototypeOf(this).constructor);
    }
    return Wrapper.prototype = Object.create(t.prototype, {
      constructor: {
        value: Wrapper,
        enumerable: !1,
        writable: !0,
        configurable: !0
      }
    }), _setPrototypeOf(Wrapper, t);
  }, _wrapNativeSuper(t);
}

var _assign = function __assign() {
  _assign = Object.assign || function __assign(t) {
    for (var s, i = 1, n = arguments.length; i < n; i++) {
      s = arguments[i];
      for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p)) t[p] = s[p];
    }
    return t;
  };
  return _assign.apply(this, arguments);
};
typeof SuppressedError === "function" ? SuppressedError : function (error, suppressed, message) {
  var e = new Error(message);
  return e.name = "SuppressedError", e.error = error, e.suppressed = suppressed, e;
};

var Chain;
(function (Chain) {
  Chain["Avalanche"] = "Avalanche";
})(Chain || (Chain = {}));

/* Do NOT modify this file; see /src.ts/_admin/update-version.ts */
/**
 *  The current version of Ethers.
 */
var version = "6.13.2";

/**
 *  Property helper functions.
 *
 *  @_subsection api/utils:Properties  [about-properties]
 */
function checkType(value, type, name) {
  var types = type.split("|").map(function (t) {
    return t.trim();
  });
  for (var i = 0; i < types.length; i++) {
    switch (type) {
      case "any":
        return;
      case "bigint":
      case "boolean":
      case "number":
      case "string":
        if (_typeof(value) === type) {
          return;
        }
    }
  }
  var error = new Error("invalid value for type ".concat(type));
  error.code = "INVALID_ARGUMENT";
  error.argument = "value.".concat(name);
  error.value = value;
  throw error;
}
function defineProperties(target, values, types) {
  for (var key in values) {
    var value = values[key];
    var type = null;
    if (type) {
      checkType(value, type, key);
    }
    Object.defineProperty(target, key, {
      enumerable: true,
      value: value,
      writable: false
    });
  }
}

function stringify(value) {
  if (value == null) {
    return "null";
  }
  if (Array.isArray(value)) {
    return "[ " + value.map(stringify).join(", ") + " ]";
  }
  if (value instanceof Uint8Array) {
    var HEX = "0123456789abcdef";
    var result = "0x";
    for (var i = 0; i < value.length; i++) {
      result += HEX[value[i] >> 4];
      result += HEX[value[i] & 0xf];
    }
    return result;
  }
  if (_typeof(value) === "object" && typeof value.toJSON === "function") {
    return stringify(value.toJSON());
  }
  switch (_typeof(value)) {
    case "boolean":
    case "symbol":
      return value.toString();
    case "bigint":
      return BigInt(value).toString();
    case "number":
      return value.toString();
    case "string":
      return JSON.stringify(value);
    case "object":
      {
        var keys = Object.keys(value);
        keys.sort();
        return "{ " + keys.map(function (k) {
          return "".concat(stringify(k), ": ").concat(stringify(value[k]));
        }).join(", ") + " }";
      }
  }
  return "[ COULD NOT SERIALIZE ]";
}
/**
 *  Returns a new Error configured to the format ethers emits errors, with
 *  the %%message%%, [[api:ErrorCode]] %%code%% and additional properties
 *  for the corresponding EthersError.
 *
 *  Each error in ethers includes the version of ethers, a
 *  machine-readable [[ErrorCode]], and depending on %%code%%, additional
 *  required properties. The error message will also include the %%message%%,
 *  ethers version, %%code%% and all additional properties, serialized.
 */
function makeError(message, code, info) {
  var shortMessage = message;
  {
    var details = [];
    if (info) {
      if ("message" in info || "code" in info || "name" in info) {
        throw new Error("value will overwrite populated values: ".concat(stringify(info)));
      }
      for (var key in info) {
        if (key === "shortMessage") {
          continue;
        }
        var value = info[key];
        //                try {
        details.push(key + "=" + stringify(value));
        //                } catch (error: any) {
        //                console.log("MMM", error.message);
        //                    details.push(key + "=[could not serialize object]");
        //                }
      }
    }
    details.push("code=".concat(code));
    details.push("version=".concat(version));
    if (details.length) {
      message += " (" + details.join(", ") + ")";
    }
  }
  var error;
  switch (code) {
    case "INVALID_ARGUMENT":
      error = new TypeError(message);
      break;
    case "NUMERIC_FAULT":
    case "BUFFER_OVERRUN":
      error = new RangeError(message);
      break;
    default:
      error = new Error(message);
  }
  defineProperties(error, {
    code: code
  });
  if (info) {
    Object.assign(error, info);
  }
  if (error.shortMessage == null) {
    defineProperties(error, {
      shortMessage: shortMessage
    });
  }
  return error;
}
/**
 *  Throws an EthersError with %%message%%, %%code%% and additional error
 *  %%info%% when %%check%% is falsish..
 *
 *  @see [[api:makeError]]
 */
function assert(check, message, code, info) {
  if (!check) {
    throw makeError(message, code, info);
  }
}
/**
 *  A simple helper to simply ensuring provided arguments match expected
 *  constraints, throwing if not.
 *
 *  In TypeScript environments, the %%check%% has been asserted true, so
 *  any further code does not need additional compile-time checks.
 */
function assertArgument(check, message, name, value) {
  assert(check, message, "INVALID_ARGUMENT", {
    argument: name,
    value: value
  });
}
["NFD", "NFC", "NFKD", "NFKC"].reduce(function (accum, form) {
  try {
    // General test for normalize
    /* c8 ignore start */
    if ("test".normalize(form) !== "test") {
      throw new Error("bad");
    }
    ;
    /* c8 ignore stop */
    if (form === "NFD") {
      var check = String.fromCharCode(0xe9).normalize("NFD");
      var expected = String.fromCharCode(0x65, 0x0301);
      /* c8 ignore start */
      if (check !== expected) {
        throw new Error("broken");
      }
      /* c8 ignore stop */
    }
    accum.push(form);
  } catch (error) {}
  return accum;
}, []);
/**
 *  Many classes use file-scoped values to guard the constructor,
 *  making it effectively private. This facilitates that pattern
 *  by ensuring the %%givenGaurd%% matches the file-scoped %%guard%%,
 *  throwing if not, indicating the %%className%% if provided.
 */
function assertPrivate(givenGuard, guard, className) {
  if (givenGuard !== guard) {
    var method = className,
      operation = "new";
    {
      method += ".";
      operation += " " + className;
    }
    assert(false, "private constructor; use ".concat(method, "from* methods"), "UNSUPPORTED_OPERATION", {
      operation: operation
    });
  }
}

/**
 *  Some data helpers.
 *
 *
 *  @_subsection api/utils:Data Helpers  [about-data]
 */
function _getBytes(value, name, copy) {
  if (value instanceof Uint8Array) {
    if (copy) {
      return new Uint8Array(value);
    }
    return value;
  }
  if (typeof value === "string" && value.match(/^0x(?:[0-9a-f][0-9a-f])*$/i)) {
    var result = new Uint8Array((value.length - 2) / 2);
    var offset = 2;
    for (var i = 0; i < result.length; i++) {
      result[i] = parseInt(value.substring(offset, offset + 2), 16);
      offset += 2;
    }
    return result;
  }
  assertArgument(false, "invalid BytesLike value", name || "value", value);
}
/**
 *  Get a typed Uint8Array for %%value%%. If already a Uint8Array
 *  the original %%value%% is returned; if a copy is required use
 *  [[getBytesCopy]].
 *
 *  @see: getBytesCopy
 */
function getBytes(value, name) {
  return _getBytes(value, name, false);
}
/**
 *  Get a typed Uint8Array for %%value%%, creating a copy if necessary
 *  to prevent any modifications of the returned value from being
 *  reflected elsewhere.
 *
 *  @see: getBytes
 */
function getBytesCopy(value, name) {
  return _getBytes(value, name, true);
}
/**
 *  Returns true if %%value%% is a valid [[HexString]].
 *
 *  If %%length%% is ``true`` or a //number//, it also checks that
 *  %%value%% is a valid [[DataHexString]] of %%length%% (if a //number//)
 *  bytes of data (e.g. ``0x1234`` is 2 bytes).
 */
function isHexString(value, length) {
  if (typeof value !== "string" || !value.match(/^0x[0-9A-Fa-f]*$/)) {
    return false;
  }
  if (typeof length === "number" && value.length !== 2 + 2 * length) {
    return false;
  }
  if (length === true && value.length % 2 !== 0) {
    return false;
  }
  return true;
}
var HexCharacters = "0123456789abcdef";
/**
 *  Returns a [[DataHexString]] representation of %%data%%.
 */
function hexlify(data) {
  var bytes = getBytes(data);
  var result = "0x";
  for (var i = 0; i < bytes.length; i++) {
    var v = bytes[i];
    result += HexCharacters[(v & 0xf0) >> 4] + HexCharacters[v & 0x0f];
  }
  return result;
}
/**
 *  Returns a [[DataHexString]] by concatenating all values
 *  within %%data%%.
 */
function concat(datas) {
  return "0x" + datas.map(function (d) {
    return hexlify(d).substring(2);
  }).join("");
}
/**
 *  Returns the length of %%data%%, in bytes.
 */
function dataLength(data) {
  if (isHexString(data, true)) {
    return (data.length - 2) / 2;
  }
  return getBytes(data).length;
}
function zeroPad(data, length, left) {
  var bytes = getBytes(data);
  assert(length >= bytes.length, "padding exceeds data length", "BUFFER_OVERRUN", {
    buffer: new Uint8Array(bytes),
    length: length,
    offset: length + 1
  });
  var result = new Uint8Array(length);
  result.fill(0);
  {
    result.set(bytes, length - bytes.length);
  }
  return hexlify(result);
}
/**
 *  Return the [[DataHexString]] of %%data%% padded on the **left**
 *  to %%length%% bytes.
 *
 *  If %%data%% already exceeds %%length%%, a [[BufferOverrunError]] is
 *  thrown.
 *
 *  This pads data the same as **values** are in Solidity
 *  (e.g. ``uint128``).
 */
function zeroPadValue(data, length) {
  return zeroPad(data, length);
}

var BN_0$2 = BigInt(0);
BigInt(1);
//const BN_Max256 = (BN_1 << BigInt(256)) - BN_1;
// IEEE 754 support 53-bits of mantissa
var maxValue = 0x1fffffffffffff;
/**
 *  Gets a BigInt from %%value%%. If it is an invalid value for
 *  a BigInt, then an ArgumentError will be thrown for %%name%%.
 */
function getBigInt(value, name) {
  switch (_typeof(value)) {
    case "bigint":
      return value;
    case "number":
      assertArgument(Number.isInteger(value), "underflow", name || "value", value);
      assertArgument(value >= -maxValue && value <= maxValue, "overflow", name || "value", value);
      return BigInt(value);
    case "string":
      try {
        if (value === "") {
          throw new Error("empty string");
        }
        if (value[0] === "-" && value[1] !== "-") {
          return -BigInt(value.substring(1));
        }
        return BigInt(value);
      } catch (e) {
        assertArgument(false, "invalid BigNumberish string: ".concat(e.message), name || "value", value);
      }
  }
  assertArgument(false, "invalid BigNumberish value", name || "value", value);
}
/**
 *  Returns %%value%% as a bigint, validating it is valid as a bigint
 *  value and that it is positive.
 */
function getUint(value, name) {
  var result = getBigInt(value, name);
  assert(result >= BN_0$2, "unsigned value cannot be negative", "NUMERIC_FAULT", {
    fault: "overflow",
    operation: "getUint",
    value: value
  });
  return result;
}
/**
 *  Gets a //number// from %%value%%. If it is an invalid value for
 *  a //number//, then an ArgumentError will be thrown for %%name%%.
 */
function getNumber(value, name) {
  switch (_typeof(value)) {
    case "bigint":
      assertArgument(value >= -maxValue && value <= maxValue, "overflow", name || "value", value);
      return Number(value);
    case "number":
      assertArgument(Number.isInteger(value), "underflow", name || "value", value);
      assertArgument(value >= -maxValue && value <= maxValue, "overflow", name || "value", value);
      return value;
    case "string":
      try {
        if (value === "") {
          throw new Error("empty string");
        }
        return getNumber(BigInt(value), name);
      } catch (e) {
        assertArgument(false, "invalid numeric string: ".concat(e.message), name || "value", value);
      }
  }
  assertArgument(false, "invalid numeric value", name || "value", value);
}
/**
 *  Converts %%value%% to a Big Endian hexstring, optionally padded to
 *  %%width%% bytes.
 */
function toBeHex(_value, _width) {
  var value = getUint(_value, "value");
  var result = value.toString(16);
  {
    var width = getNumber(_width, "width");
    assert(width * 2 >= result.length, "value exceeds width (".concat(width, " bytes)"), "NUMERIC_FAULT", {
      operation: "toBeHex",
      fault: "overflow",
      value: _value
    });
    // Pad the value to the required width
    while (result.length < width * 2) {
      result = "0" + result;
    }
  }
  return "0x" + result;
}
/**
 *  Converts %%value%% to a Big Endian Uint8Array.
 */
function toBeArray(_value) {
  var value = getUint(_value, "value");
  if (value === BN_0$2) {
    return new Uint8Array([]);
  }
  var hex = value.toString(16);
  if (hex.length % 2) {
    hex = "0" + hex;
  }
  var result = new Uint8Array(hex.length / 2);
  for (var i = 0; i < result.length; i++) {
    var offset = i * 2;
    result[i] = parseInt(hex.substring(offset, offset + 2), 16);
  }
  return result;
}

/**
 *  Using strings in Ethereum (or any security-basd system) requires
 *  additional care. These utilities attempt to mitigate some of the
 *  safety issues as well as provide the ability to recover and analyse
 *  strings.
 *
 *  @_subsection api/utils:Strings and UTF-8  [about-strings]
 */
// http://stackoverflow.com/questions/18729405/how-to-convert-utf8-string-to-byte-array
/**
 *  Returns the UTF-8 byte representation of %%str%%.
 *
 *  If %%form%% is specified, the string is normalized.
 */
function toUtf8Bytes(str, form) {
  assertArgument(typeof str === "string", "invalid string value", "str", str);
  var result = [];
  for (var i = 0; i < str.length; i++) {
    var c = str.charCodeAt(i);
    if (c < 0x80) {
      result.push(c);
    } else if (c < 0x800) {
      result.push(c >> 6 | 0xc0);
      result.push(c & 0x3f | 0x80);
    } else if ((c & 0xfc00) == 0xd800) {
      i++;
      var c2 = str.charCodeAt(i);
      assertArgument(i < str.length && (c2 & 0xfc00) === 0xdc00, "invalid surrogate pair", "str", str);
      // Surrogate Pair
      var pair = 0x10000 + ((c & 0x03ff) << 10) + (c2 & 0x03ff);
      result.push(pair >> 18 | 0xf0);
      result.push(pair >> 12 & 0x3f | 0x80);
      result.push(pair >> 6 & 0x3f | 0x80);
      result.push(pair & 0x3f | 0x80);
    } else {
      result.push(c >> 12 | 0xe0);
      result.push(c >> 6 & 0x3f | 0x80);
      result.push(c & 0x3f | 0x80);
    }
  }
  return new Uint8Array(result);
}

function number(n) {
  if (!Number.isSafeInteger(n) || n < 0) throw new Error("Wrong positive integer: ".concat(n));
}
function bytes(b) {
  if (!(b instanceof Uint8Array)) throw new Error('Expected Uint8Array');
  for (var _len = arguments.length, lengths = new Array(_len > 1 ? _len - 1 : 0), _key = 1; _key < _len; _key++) {
    lengths[_key - 1] = arguments[_key];
  }
  if (lengths.length > 0 && !lengths.includes(b.length)) throw new Error("Expected Uint8Array of length ".concat(lengths, ", not of length=").concat(b.length));
}
function hash(hash) {
  if (typeof hash !== 'function' || typeof hash.create !== 'function') throw new Error('Hash should be wrapped by utils.wrapConstructor');
  number(hash.outputLen);
  number(hash.blockLen);
}
function exists(instance) {
  var checkFinished = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : true;
  if (instance.destroyed) throw new Error('Hash instance has been destroyed');
  if (checkFinished && instance.finished) throw new Error('Hash#digest() has already been called');
}
function output(out, instance) {
  bytes(out);
  var min = instance.outputLen;
  if (out.length < min) {
    throw new Error("digestInto() expects output buffer of length at least ".concat(min));
  }
}

var U32_MASK64 = /* @__PURE__ */BigInt(Math.pow(2, 32) - 1);
var _32n = /* @__PURE__ */BigInt(32);
// We are not using BigUint64Array, because they are extremely slow as per 2022
function fromBig(n) {
  var le = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : false;
  if (le) return {
    h: Number(n & U32_MASK64),
    l: Number(n >> _32n & U32_MASK64)
  };
  return {
    h: Number(n >> _32n & U32_MASK64) | 0,
    l: Number(n & U32_MASK64) | 0
  };
}
function split(lst) {
  var le = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : false;
  var Ah = new Uint32Array(lst.length);
  var Al = new Uint32Array(lst.length);
  for (var i = 0; i < lst.length; i++) {
    var _fromBig = fromBig(lst[i], le),
      h = _fromBig.h,
      l = _fromBig.l;
    var _ref = [h, l];
    Ah[i] = _ref[0];
    Al[i] = _ref[1];
  }
  return [Ah, Al];
}
// Left rotate for Shift in [1, 32)
var rotlSH = function rotlSH(h, l, s) {
  return h << s | l >>> 32 - s;
};
var rotlSL = function rotlSL(h, l, s) {
  return l << s | h >>> 32 - s;
};
// Left rotate for Shift in (32, 64), NOTE: 32 is special case.
var rotlBH = function rotlBH(h, l, s) {
  return l << s - 32 | h >>> 64 - s;
};
var rotlBL = function rotlBL(h, l, s) {
  return h << s - 32 | l >>> 64 - s;
};

var crypto = (typeof globalThis === "undefined" ? "undefined" : _typeof(globalThis)) === 'object' && 'crypto' in globalThis ? globalThis.crypto : undefined;

var u8a$1 = function u8a(a) {
  return a instanceof Uint8Array;
};
var u32 = function u32(arr) {
  return new Uint32Array(arr.buffer, arr.byteOffset, Math.floor(arr.byteLength / 4));
};
// Cast array to view
var createView = function createView(arr) {
  return new DataView(arr.buffer, arr.byteOffset, arr.byteLength);
};
// The rotate right (circular right shift) operation for uint32
var rotr = function rotr(word, shift) {
  return word << 32 - shift | word >>> shift;
};
// big-endian hardware is rare. Just in case someone still decides to run hashes:
// early-throw an error because we don't support BE yet.
var isLE = new Uint8Array(new Uint32Array([0x11223344]).buffer)[0] === 0x44;
if (!isLE) throw new Error('Non little-endian hardware is not supported');
function utf8ToBytes$1(str) {
  if (typeof str !== 'string') throw new Error("utf8ToBytes expected string, got ".concat(_typeof(str)));
  return new Uint8Array(new TextEncoder().encode(str)); // https://bugzil.la/1681809
}
/**
 * Normalizes (non-hex) string or Uint8Array to Uint8Array.
 * Warning: when Uint8Array is passed, it would NOT get copied.
 * Keep in mind for future mutable operations.
 */
function toBytes(data) {
  if (typeof data === 'string') data = utf8ToBytes$1(data);
  if (!u8a$1(data)) throw new Error("expected Uint8Array, got ".concat(_typeof(data)));
  return data;
}
/**
 * Copies several Uint8Arrays into one.
 */
function concatBytes$1() {
  for (var _len = arguments.length, arrays = new Array(_len), _key = 0; _key < _len; _key++) {
    arrays[_key] = arguments[_key];
  }
  var r = new Uint8Array(arrays.reduce(function (sum, a) {
    return sum + a.length;
  }, 0));
  var pad = 0; // walk through each item, ensure they have proper type
  arrays.forEach(function (a) {
    if (!u8a$1(a)) throw new Error('Uint8Array expected');
    r.set(a, pad);
    pad += a.length;
  });
  return r;
}
// For runtime check if class implements interface
var Hash = /*#__PURE__*/function () {
  function Hash() {
    _classCallCheck(this, Hash);
  }
  return _createClass(Hash, [{
    key: "clone",
    value:
    // Safe version that clones internal state
    function clone() {
      return this._cloneInto();
    }
  }]);
}();
function wrapConstructor(hashCons) {
  var hashC = function hashC(msg) {
    return hashCons().update(toBytes(msg)).digest();
  };
  var tmp = hashCons();
  hashC.outputLen = tmp.outputLen;
  hashC.blockLen = tmp.blockLen;
  hashC.create = function () {
    return hashCons();
  };
  return hashC;
}
/**
 * Secure PRNG. Uses `crypto.getRandomValues`, which defers to OS.
 */
function randomBytes() {
  var bytesLength = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : 32;
  if (crypto && typeof crypto.getRandomValues === 'function') {
    return crypto.getRandomValues(new Uint8Array(bytesLength));
  }
  throw new Error('crypto.getRandomValues must be defined');
}

// SHA3 (keccak) is based on a new design: basically, the internal state is bigger than output size.
// It's called a sponge function.
// Various per round constants calculations
var SHA3_PI = [],
  SHA3_ROTL = [],
  _SHA3_IOTA = [];
var _0n$4 = /* @__PURE__ */BigInt(0);
var _1n$5 = /* @__PURE__ */BigInt(1);
var _2n$3 = /* @__PURE__ */BigInt(2);
var _7n = /* @__PURE__ */BigInt(7);
var _256n = /* @__PURE__ */BigInt(256);
var _0x71n = /* @__PURE__ */BigInt(0x71);
for (var round = 0, R = _1n$5, x = 1, y = 0; round < 24; round++) {
  // Pi
  var _ref = [y, (2 * x + 3 * y) % 5];
  x = _ref[0];
  y = _ref[1];
  SHA3_PI.push(2 * (5 * y + x));
  // Rotational
  SHA3_ROTL.push((round + 1) * (round + 2) / 2 % 64);
  // Iota
  var t = _0n$4;
  for (var j = 0; j < 7; j++) {
    R = (R << _1n$5 ^ (R >> _7n) * _0x71n) % _256n;
    if (R & _2n$3) t ^= _1n$5 << (_1n$5 << /* @__PURE__ */BigInt(j)) - _1n$5;
  }
  _SHA3_IOTA.push(t);
}
var _split = /* @__PURE__ */split(_SHA3_IOTA, true),
  _split2 = _slicedToArray(_split, 2),
  SHA3_IOTA_H = _split2[0],
  SHA3_IOTA_L = _split2[1];
// Left rotation (without 0, 32, 64)
var rotlH = function rotlH(h, l, s) {
  return s > 32 ? rotlBH(h, l, s) : rotlSH(h, l, s);
};
var rotlL = function rotlL(h, l, s) {
  return s > 32 ? rotlBL(h, l, s) : rotlSL(h, l, s);
};
// Same as keccakf1600, but allows to skip some rounds
function keccakP(s) {
  var rounds = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 24;
  var B = new Uint32Array(5 * 2);
  // NOTE: all indices are x2 since we store state as u32 instead of u64 (bigints to slow in js)
  for (var _round = 24 - rounds; _round < 24; _round++) {
    // Theta θ
    for (var _x = 0; _x < 10; _x++) B[_x] = s[_x] ^ s[_x + 10] ^ s[_x + 20] ^ s[_x + 30] ^ s[_x + 40];
    for (var _x2 = 0; _x2 < 10; _x2 += 2) {
      var idx1 = (_x2 + 8) % 10;
      var idx0 = (_x2 + 2) % 10;
      var B0 = B[idx0];
      var B1 = B[idx0 + 1];
      var Th = rotlH(B0, B1, 1) ^ B[idx1];
      var Tl = rotlL(B0, B1, 1) ^ B[idx1 + 1];
      for (var _y = 0; _y < 50; _y += 10) {
        s[_x2 + _y] ^= Th;
        s[_x2 + _y + 1] ^= Tl;
      }
    }
    // Rho (ρ) and Pi (π)
    var curH = s[2];
    var curL = s[3];
    for (var _t = 0; _t < 24; _t++) {
      var shift = SHA3_ROTL[_t];
      var _Th = rotlH(curH, curL, shift);
      var _Tl = rotlL(curH, curL, shift);
      var PI = SHA3_PI[_t];
      curH = s[PI];
      curL = s[PI + 1];
      s[PI] = _Th;
      s[PI + 1] = _Tl;
    }
    // Chi (χ)
    for (var _y2 = 0; _y2 < 50; _y2 += 10) {
      for (var _x3 = 0; _x3 < 10; _x3++) B[_x3] = s[_y2 + _x3];
      for (var _x4 = 0; _x4 < 10; _x4++) s[_y2 + _x4] ^= ~B[(_x4 + 2) % 10] & B[(_x4 + 4) % 10];
    }
    // Iota (ι)
    s[0] ^= SHA3_IOTA_H[_round];
    s[1] ^= SHA3_IOTA_L[_round];
  }
  B.fill(0);
}
var Keccak = /*#__PURE__*/function (_Hash) {
  // NOTE: we accept arguments in bytes instead of bits here.
  function Keccak(blockLen, suffix, outputLen) {
    var _this;
    var enableXOF = arguments.length > 3 && arguments[3] !== undefined ? arguments[3] : false;
    var rounds = arguments.length > 4 && arguments[4] !== undefined ? arguments[4] : 24;
    _classCallCheck(this, Keccak);
    _this = _callSuper(this, Keccak);
    _this.blockLen = blockLen;
    _this.suffix = suffix;
    _this.outputLen = outputLen;
    _this.enableXOF = enableXOF;
    _this.rounds = rounds;
    _this.pos = 0;
    _this.posOut = 0;
    _this.finished = false;
    _this.destroyed = false;
    // Can be passed from user as dkLen
    number(outputLen);
    // 1600 = 5x5 matrix of 64bit.  1600 bits === 200 bytes
    if (0 >= _this.blockLen || _this.blockLen >= 200) throw new Error('Sha3 supports only keccak-f1600 function');
    _this.state = new Uint8Array(200);
    _this.state32 = u32(_this.state);
    return _this;
  }
  _inherits(Keccak, _Hash);
  return _createClass(Keccak, [{
    key: "keccak",
    value: function keccak() {
      keccakP(this.state32, this.rounds);
      this.posOut = 0;
      this.pos = 0;
    }
  }, {
    key: "update",
    value: function update(data) {
      exists(this);
      var blockLen = this.blockLen,
        state = this.state;
      data = toBytes(data);
      var len = data.length;
      for (var pos = 0; pos < len;) {
        var take = Math.min(blockLen - this.pos, len - pos);
        for (var i = 0; i < take; i++) state[this.pos++] ^= data[pos++];
        if (this.pos === blockLen) this.keccak();
      }
      return this;
    }
  }, {
    key: "finish",
    value: function finish() {
      if (this.finished) return;
      this.finished = true;
      var state = this.state,
        suffix = this.suffix,
        pos = this.pos,
        blockLen = this.blockLen;
      // Do the padding
      state[pos] ^= suffix;
      if ((suffix & 0x80) !== 0 && pos === blockLen - 1) this.keccak();
      state[blockLen - 1] ^= 0x80;
      this.keccak();
    }
  }, {
    key: "writeInto",
    value: function writeInto(out) {
      exists(this, false);
      bytes(out);
      this.finish();
      var bufferOut = this.state;
      var blockLen = this.blockLen;
      for (var pos = 0, len = out.length; pos < len;) {
        if (this.posOut >= blockLen) this.keccak();
        var take = Math.min(blockLen - this.posOut, len - pos);
        out.set(bufferOut.subarray(this.posOut, this.posOut + take), pos);
        this.posOut += take;
        pos += take;
      }
      return out;
    }
  }, {
    key: "xofInto",
    value: function xofInto(out) {
      // Sha3/Keccak usage with XOF is probably mistake, only SHAKE instances can do XOF
      if (!this.enableXOF) throw new Error('XOF is not possible for this instance');
      return this.writeInto(out);
    }
  }, {
    key: "xof",
    value: function xof(bytes) {
      number(bytes);
      return this.xofInto(new Uint8Array(bytes));
    }
  }, {
    key: "digestInto",
    value: function digestInto(out) {
      output(out, this);
      if (this.finished) throw new Error('digest() was already called');
      this.writeInto(out);
      this.destroy();
      return out;
    }
  }, {
    key: "digest",
    value: function digest() {
      return this.digestInto(new Uint8Array(this.outputLen));
    }
  }, {
    key: "destroy",
    value: function destroy() {
      this.destroyed = true;
      this.state.fill(0);
    }
  }, {
    key: "_cloneInto",
    value: function _cloneInto(to) {
      var blockLen = this.blockLen,
        suffix = this.suffix,
        outputLen = this.outputLen,
        rounds = this.rounds,
        enableXOF = this.enableXOF;
      to || (to = new Keccak(blockLen, suffix, outputLen, enableXOF, rounds));
      to.state32.set(this.state32);
      to.pos = this.pos;
      to.posOut = this.posOut;
      to.finished = this.finished;
      to.rounds = rounds;
      // Suffix can change in cSHAKE
      to.suffix = suffix;
      to.outputLen = outputLen;
      to.enableXOF = enableXOF;
      to.destroyed = this.destroyed;
      return to;
    }
  }]);
}(Hash);
var gen = function gen(suffix, blockLen, outputLen) {
  return wrapConstructor(function () {
    return new Keccak(blockLen, suffix, outputLen);
  });
};
/**
 * keccak-256 hash function. Different from SHA3-256.
 * @param message - that would be hashed
 */
var keccak_256 = /* @__PURE__ */gen(0x01, 136, 256 / 8);

/**
 *  Cryptographic hashing functions
 *
 *  @_subsection: api/crypto:Hash Functions [about-crypto-hashing]
 */
var locked = false;
var _keccak256 = function _keccak256(data) {
  return keccak_256(data);
};
var __keccak256 = _keccak256;
/**
 *  Compute the cryptographic KECCAK256 hash of %%data%%.
 *
 *  The %%data%% **must** be a data representation, to compute the
 *  hash of UTF-8 data use the [[id]] function.
 *
 *  @returns DataHexstring
 *  @example:
 *    keccak256("0x")
 *    //_result:
 *
 *    keccak256("0x1337")
 *    //_result:
 *
 *    keccak256(new Uint8Array([ 0x13, 0x37 ]))
 *    //_result:
 *
 *    // Strings are assumed to be DataHexString, otherwise it will
 *    // throw. To hash UTF-8 data, see the note above.
 *    keccak256("Hello World")
 *    //_error:
 */
function keccak256(_data) {
  var data = getBytes(_data, "data");
  return hexlify(__keccak256(data));
}
keccak256._ = _keccak256;
keccak256.lock = function () {
  locked = true;
};
keccak256.register = function (func) {
  if (locked) {
    throw new TypeError("keccak256 is locked");
  }
  __keccak256 = func;
};
Object.freeze(keccak256);

// Polyfill for Safari 14
function setBigUint64(view, byteOffset, value, isLE) {
  if (typeof view.setBigUint64 === 'function') return view.setBigUint64(byteOffset, value, isLE);
  var _32n = BigInt(32);
  var _u32_max = BigInt(0xffffffff);
  var wh = Number(value >> _32n & _u32_max);
  var wl = Number(value & _u32_max);
  var h = isLE ? 4 : 0;
  var l = isLE ? 0 : 4;
  view.setUint32(byteOffset + h, wh, isLE);
  view.setUint32(byteOffset + l, wl, isLE);
}
// Base SHA2 class (RFC 6234)
var SHA2 = /*#__PURE__*/function (_Hash) {
  function SHA2(blockLen, outputLen, padOffset, isLE) {
    var _this;
    _classCallCheck(this, SHA2);
    _this = _callSuper(this, SHA2);
    _this.blockLen = blockLen;
    _this.outputLen = outputLen;
    _this.padOffset = padOffset;
    _this.isLE = isLE;
    _this.finished = false;
    _this.length = 0;
    _this.pos = 0;
    _this.destroyed = false;
    _this.buffer = new Uint8Array(blockLen);
    _this.view = createView(_this.buffer);
    return _this;
  }
  _inherits(SHA2, _Hash);
  return _createClass(SHA2, [{
    key: "update",
    value: function update(data) {
      exists(this);
      var view = this.view,
        buffer = this.buffer,
        blockLen = this.blockLen;
      data = toBytes(data);
      var len = data.length;
      for (var pos = 0; pos < len;) {
        var take = Math.min(blockLen - this.pos, len - pos);
        // Fast path: we have at least one block in input, cast it to view and process
        if (take === blockLen) {
          var dataView = createView(data);
          for (; blockLen <= len - pos; pos += blockLen) this.process(dataView, pos);
          continue;
        }
        buffer.set(data.subarray(pos, pos + take), this.pos);
        this.pos += take;
        pos += take;
        if (this.pos === blockLen) {
          this.process(view, 0);
          this.pos = 0;
        }
      }
      this.length += data.length;
      this.roundClean();
      return this;
    }
  }, {
    key: "digestInto",
    value: function digestInto(out) {
      exists(this);
      output(out, this);
      this.finished = true;
      // Padding
      // We can avoid allocation of buffer for padding completely if it
      // was previously not allocated here. But it won't change performance.
      var buffer = this.buffer,
        view = this.view,
        blockLen = this.blockLen,
        isLE = this.isLE;
      var pos = this.pos;
      // append the bit '1' to the message
      buffer[pos++] = 128;
      this.buffer.subarray(pos).fill(0);
      // we have less than padOffset left in buffer, so we cannot put length in current block, need process it and pad again
      if (this.padOffset > blockLen - pos) {
        this.process(view, 0);
        pos = 0;
      }
      // Pad until full block byte with zeros
      for (var i = pos; i < blockLen; i++) buffer[i] = 0;
      // Note: sha512 requires length to be 128bit integer, but length in JS will overflow before that
      // You need to write around 2 exabytes (u64_max / 8 / (1024**6)) for this to happen.
      // So we just write lowest 64 bits of that value.
      setBigUint64(view, blockLen - 8, BigInt(this.length * 8), isLE);
      this.process(view, 0);
      var oview = createView(out);
      var len = this.outputLen;
      // NOTE: we do division by 4 later, which should be fused in single op with modulo by JIT
      if (len % 4) throw new Error('_sha2: outputLen should be aligned to 32bit');
      var outLen = len / 4;
      var state = this.get();
      if (outLen > state.length) throw new Error('_sha2: outputLen bigger than state');
      for (var _i = 0; _i < outLen; _i++) oview.setUint32(4 * _i, state[_i], isLE);
    }
  }, {
    key: "digest",
    value: function digest() {
      var buffer = this.buffer,
        outputLen = this.outputLen;
      this.digestInto(buffer);
      var res = buffer.slice(0, outputLen);
      this.destroy();
      return res;
    }
  }, {
    key: "_cloneInto",
    value: function _cloneInto(to) {
      var _to;
      to || (to = new this.constructor());
      (_to = to).set.apply(_to, _toConsumableArray(this.get()));
      var blockLen = this.blockLen,
        buffer = this.buffer,
        length = this.length,
        finished = this.finished,
        destroyed = this.destroyed,
        pos = this.pos;
      to.length = length;
      to.pos = pos;
      to.finished = finished;
      to.destroyed = destroyed;
      if (length % blockLen) to.buffer.set(buffer);
      return to;
    }
  }]);
}(Hash);

// SHA2-256 need to try 2^128 hashes to execute birthday attack.
// BTC network is doing 2^67 hashes/sec as per early 2023.
// Choice: a ? b : c
var Chi = function Chi(a, b, c) {
  return a & b ^ ~a & c;
};
// Majority function, true if any two inpust is true
var Maj = function Maj(a, b, c) {
  return a & b ^ a & c ^ b & c;
};
// Round constants:
// first 32 bits of the fractional parts of the cube roots of the first 64 primes 2..311)
// prettier-ignore
var SHA256_K = /* @__PURE__ */new Uint32Array([0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5, 0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174, 0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da, 0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967, 0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85, 0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070, 0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3, 0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2]);
// Initial state (first 32 bits of the fractional parts of the square roots of the first 8 primes 2..19):
// prettier-ignore
var IV = /* @__PURE__ */new Uint32Array([0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a, 0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19]);
// Temporary buffer, not used to store anything between runs
// Named this way because it matches specification.
var SHA256_W = /* @__PURE__ */new Uint32Array(64);
var SHA256 = /*#__PURE__*/function (_SHA) {
  function SHA256() {
    var _this;
    _classCallCheck(this, SHA256);
    _this = _callSuper(this, SHA256, [64, 32, 8, false]);
    // We cannot use array here since array allows indexing by variable
    // which means optimizer/compiler cannot use registers.
    _this.A = IV[0] | 0;
    _this.B = IV[1] | 0;
    _this.C = IV[2] | 0;
    _this.D = IV[3] | 0;
    _this.E = IV[4] | 0;
    _this.F = IV[5] | 0;
    _this.G = IV[6] | 0;
    _this.H = IV[7] | 0;
    return _this;
  }
  _inherits(SHA256, _SHA);
  return _createClass(SHA256, [{
    key: "get",
    value: function get() {
      var A = this.A,
        B = this.B,
        C = this.C,
        D = this.D,
        E = this.E,
        F = this.F,
        G = this.G,
        H = this.H;
      return [A, B, C, D, E, F, G, H];
    }
    // prettier-ignore
  }, {
    key: "set",
    value: function set(A, B, C, D, E, F, G, H) {
      this.A = A | 0;
      this.B = B | 0;
      this.C = C | 0;
      this.D = D | 0;
      this.E = E | 0;
      this.F = F | 0;
      this.G = G | 0;
      this.H = H | 0;
    }
  }, {
    key: "process",
    value: function process(view, offset) {
      // Extend the first 16 words into the remaining 48 words w[16..63] of the message schedule array
      for (var i = 0; i < 16; i++, offset += 4) SHA256_W[i] = view.getUint32(offset, false);
      for (var _i = 16; _i < 64; _i++) {
        var W15 = SHA256_W[_i - 15];
        var W2 = SHA256_W[_i - 2];
        var s0 = rotr(W15, 7) ^ rotr(W15, 18) ^ W15 >>> 3;
        var s1 = rotr(W2, 17) ^ rotr(W2, 19) ^ W2 >>> 10;
        SHA256_W[_i] = s1 + SHA256_W[_i - 7] + s0 + SHA256_W[_i - 16] | 0;
      }
      // Compression function main loop, 64 rounds
      var A = this.A,
        B = this.B,
        C = this.C,
        D = this.D,
        E = this.E,
        F = this.F,
        G = this.G,
        H = this.H;
      for (var _i2 = 0; _i2 < 64; _i2++) {
        var sigma1 = rotr(E, 6) ^ rotr(E, 11) ^ rotr(E, 25);
        var T1 = H + sigma1 + Chi(E, F, G) + SHA256_K[_i2] + SHA256_W[_i2] | 0;
        var sigma0 = rotr(A, 2) ^ rotr(A, 13) ^ rotr(A, 22);
        var T2 = sigma0 + Maj(A, B, C) | 0;
        H = G;
        G = F;
        F = E;
        E = D + T1 | 0;
        D = C;
        C = B;
        B = A;
        A = T1 + T2 | 0;
      }
      // Add the compressed chunk to the current hash value
      A = A + this.A | 0;
      B = B + this.B | 0;
      C = C + this.C | 0;
      D = D + this.D | 0;
      E = E + this.E | 0;
      F = F + this.F | 0;
      G = G + this.G | 0;
      H = H + this.H | 0;
      this.set(A, B, C, D, E, F, G, H);
    }
  }, {
    key: "roundClean",
    value: function roundClean() {
      SHA256_W.fill(0);
    }
  }, {
    key: "destroy",
    value: function destroy() {
      this.set(0, 0, 0, 0, 0, 0, 0, 0);
      this.buffer.fill(0);
    }
  }]);
}(SHA2); // Constants from https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.180-4.pdf
/**
 * SHA2-256 hash function
 * @param message - data that would be hashed
 */
var sha256 = /* @__PURE__ */wrapConstructor(function () {
  return new SHA256();
});

// HMAC (RFC 2104)
var HMAC = /*#__PURE__*/function (_Hash) {
  function HMAC(hash$1, _key) {
    var _this;
    _classCallCheck(this, HMAC);
    _this = _callSuper(this, HMAC);
    _this.finished = false;
    _this.destroyed = false;
    hash(hash$1);
    var key = toBytes(_key);
    _this.iHash = hash$1.create();
    if (typeof _this.iHash.update !== 'function') throw new Error('Expected instance of class which extends utils.Hash');
    _this.blockLen = _this.iHash.blockLen;
    _this.outputLen = _this.iHash.outputLen;
    var blockLen = _this.blockLen;
    var pad = new Uint8Array(blockLen);
    // blockLen can be bigger than outputLen
    pad.set(key.length > blockLen ? hash$1.create().update(key).digest() : key);
    for (var i = 0; i < pad.length; i++) pad[i] ^= 0x36;
    _this.iHash.update(pad);
    // By doing update (processing of first block) of outer hash here we can re-use it between multiple calls via clone
    _this.oHash = hash$1.create();
    // Undo internal XOR && apply outer XOR
    for (var _i = 0; _i < pad.length; _i++) pad[_i] ^= 0x36 ^ 0x5c;
    _this.oHash.update(pad);
    pad.fill(0);
    return _this;
  }
  _inherits(HMAC, _Hash);
  return _createClass(HMAC, [{
    key: "update",
    value: function update(buf) {
      exists(this);
      this.iHash.update(buf);
      return this;
    }
  }, {
    key: "digestInto",
    value: function digestInto(out) {
      exists(this);
      bytes(out, this.outputLen);
      this.finished = true;
      this.iHash.digestInto(out);
      this.oHash.update(out);
      this.oHash.digestInto(out);
      this.destroy();
    }
  }, {
    key: "digest",
    value: function digest() {
      var out = new Uint8Array(this.oHash.outputLen);
      this.digestInto(out);
      return out;
    }
  }, {
    key: "_cloneInto",
    value: function _cloneInto(to) {
      // Create new instance without calling constructor since key already in state and we don't know it.
      to || (to = Object.create(Object.getPrototypeOf(this), {}));
      var oHash = this.oHash,
        iHash = this.iHash,
        finished = this.finished,
        destroyed = this.destroyed,
        blockLen = this.blockLen,
        outputLen = this.outputLen;
      to = to;
      to.finished = finished;
      to.destroyed = destroyed;
      to.blockLen = blockLen;
      to.outputLen = outputLen;
      to.oHash = oHash._cloneInto(to.oHash);
      to.iHash = iHash._cloneInto(to.iHash);
      return to;
    }
  }, {
    key: "destroy",
    value: function destroy() {
      this.destroyed = true;
      this.oHash.destroy();
      this.iHash.destroy();
    }
  }]);
}(Hash);
/**
 * HMAC: RFC2104 message authentication code.
 * @param hash - function that would be used e.g. sha256
 * @param key - message key
 * @param message - message data
 */
var hmac = function hmac(hash, key, message) {
  return new HMAC(hash, key).update(message).digest();
};
hmac.create = function (hash, key) {
  return new HMAC(hash, key);
};

/*! noble-curves - MIT License (c) 2022 Paul Miller (paulmillr.com) */
// 100 lines of code in the file are duplicated from noble-hashes (utils).
// This is OK: `abstract` directory does not use noble-hashes.
// User may opt-in into using different hashing library. This way, noble-hashes
// won't be included into their bundle.
var _0n$3 = BigInt(0);
var _1n$4 = BigInt(1);
var _2n$2 = BigInt(2);
var u8a = function u8a(a) {
  return a instanceof Uint8Array;
};
var hexes = /* @__PURE__ */Array.from({
  length: 256
}, function (_, i) {
  return i.toString(16).padStart(2, '0');
});
/**
 * @example bytesToHex(Uint8Array.from([0xca, 0xfe, 0x01, 0x23])) // 'cafe0123'
 */
function bytesToHex(bytes) {
  if (!u8a(bytes)) throw new Error('Uint8Array expected');
  // pre-caching improves the speed 6x
  var hex = '';
  for (var i = 0; i < bytes.length; i++) {
    hex += hexes[bytes[i]];
  }
  return hex;
}
function numberToHexUnpadded(num) {
  var hex = num.toString(16);
  return hex.length & 1 ? "0".concat(hex) : hex;
}
function hexToNumber(hex) {
  if (typeof hex !== 'string') throw new Error('hex string expected, got ' + _typeof(hex));
  // Big Endian
  return BigInt(hex === '' ? '0' : "0x".concat(hex));
}
/**
 * @example hexToBytes('cafe0123') // Uint8Array.from([0xca, 0xfe, 0x01, 0x23])
 */
function hexToBytes(hex) {
  if (typeof hex !== 'string') throw new Error('hex string expected, got ' + _typeof(hex));
  var len = hex.length;
  if (len % 2) throw new Error('padded hex string expected, got unpadded hex of length ' + len);
  var array = new Uint8Array(len / 2);
  for (var i = 0; i < array.length; i++) {
    var j = i * 2;
    var hexByte = hex.slice(j, j + 2);
    var _byte = Number.parseInt(hexByte, 16);
    if (Number.isNaN(_byte) || _byte < 0) throw new Error('Invalid byte sequence');
    array[i] = _byte;
  }
  return array;
}
// BE: Big Endian, LE: Little Endian
function bytesToNumberBE(bytes) {
  return hexToNumber(bytesToHex(bytes));
}
function bytesToNumberLE(bytes) {
  if (!u8a(bytes)) throw new Error('Uint8Array expected');
  return hexToNumber(bytesToHex(Uint8Array.from(bytes).reverse()));
}
function numberToBytesBE(n, len) {
  return hexToBytes(n.toString(16).padStart(len * 2, '0'));
}
function numberToBytesLE(n, len) {
  return numberToBytesBE(n, len).reverse();
}
// Unpadded, rarely used
function numberToVarBytesBE(n) {
  return hexToBytes(numberToHexUnpadded(n));
}
/**
 * Takes hex string or Uint8Array, converts to Uint8Array.
 * Validates output length.
 * Will throw error for other types.
 * @param title descriptive title for an error e.g. 'private key'
 * @param hex hex string or Uint8Array
 * @param expectedLength optional, will compare to result array's length
 * @returns
 */
function ensureBytes(title, hex, expectedLength) {
  var res;
  if (typeof hex === 'string') {
    try {
      res = hexToBytes(hex);
    } catch (e) {
      throw new Error("".concat(title, " must be valid hex string, got \"").concat(hex, "\". Cause: ").concat(e));
    }
  } else if (u8a(hex)) {
    // Uint8Array.from() instead of hash.slice() because node.js Buffer
    // is instance of Uint8Array, and its slice() creates **mutable** copy
    res = Uint8Array.from(hex);
  } else {
    throw new Error("".concat(title, " must be hex string or Uint8Array"));
  }
  var len = res.length;
  if (typeof expectedLength === 'number' && len !== expectedLength) throw new Error("".concat(title, " expected ").concat(expectedLength, " bytes, got ").concat(len));
  return res;
}
/**
 * Copies several Uint8Arrays into one.
 */
function concatBytes() {
  for (var _len = arguments.length, arrays = new Array(_len), _key = 0; _key < _len; _key++) {
    arrays[_key] = arguments[_key];
  }
  var r = new Uint8Array(arrays.reduce(function (sum, a) {
    return sum + a.length;
  }, 0));
  var pad = 0; // walk through each item, ensure they have proper type
  arrays.forEach(function (a) {
    if (!u8a(a)) throw new Error('Uint8Array expected');
    r.set(a, pad);
    pad += a.length;
  });
  return r;
}
function equalBytes(b1, b2) {
  // We don't care about timing attacks here
  if (b1.length !== b2.length) return false;
  for (var i = 0; i < b1.length; i++) if (b1[i] !== b2[i]) return false;
  return true;
}
/**
 * @example utf8ToBytes('abc') // new Uint8Array([97, 98, 99])
 */
function utf8ToBytes(str) {
  if (typeof str !== 'string') throw new Error("utf8ToBytes expected string, got ".concat(_typeof(str)));
  return new Uint8Array(new TextEncoder().encode(str)); // https://bugzil.la/1681809
}
// Bit operations
/**
 * Calculates amount of bits in a bigint.
 * Same as `n.toString(2).length`
 */
function bitLen(n) {
  var len;
  for (len = 0; n > _0n$3; n >>= _1n$4, len += 1);
  return len;
}
/**
 * Gets single bit at position.
 * NOTE: first bit position is 0 (same as arrays)
 * Same as `!!+Array.from(n.toString(2)).reverse()[pos]`
 */
function bitGet(n, pos) {
  return n >> BigInt(pos) & _1n$4;
}
/**
 * Sets single bit at position.
 */
var bitSet = function bitSet(n, pos, value) {
  return n | (value ? _1n$4 : _0n$3) << BigInt(pos);
};
/**
 * Calculate mask for N bits. Not using ** operator with bigints because of old engines.
 * Same as BigInt(`0b${Array(i).fill('1').join('')}`)
 */
var bitMask = function bitMask(n) {
  return (_2n$2 << BigInt(n - 1)) - _1n$4;
};
// DRBG
var u8n = function u8n(data) {
  return new Uint8Array(data);
}; // creates Uint8Array
var u8fr = function u8fr(arr) {
  return Uint8Array.from(arr);
}; // another shortcut
/**
 * Minimal HMAC-DRBG from NIST 800-90 for RFC6979 sigs.
 * @returns function that will call DRBG until 2nd arg returns something meaningful
 * @example
 *   const drbg = createHmacDRBG<Key>(32, 32, hmac);
 *   drbg(seed, bytesToKey); // bytesToKey must return Key or undefined
 */
function createHmacDrbg(hashLen, qByteLen, hmacFn) {
  if (typeof hashLen !== 'number' || hashLen < 2) throw new Error('hashLen must be a number');
  if (typeof qByteLen !== 'number' || qByteLen < 2) throw new Error('qByteLen must be a number');
  if (typeof hmacFn !== 'function') throw new Error('hmacFn must be a function');
  // Step B, Step C: set hashLen to 8*ceil(hlen/8)
  var v = u8n(hashLen); // Minimal non-full-spec HMAC-DRBG from NIST 800-90 for RFC6979 sigs.
  var k = u8n(hashLen); // Steps B and C of RFC6979 3.2: set hashLen, in our case always same
  var i = 0; // Iterations counter, will throw when over 1000
  var reset = function reset() {
    v.fill(1);
    k.fill(0);
    i = 0;
  };
  var h = function h() {
    for (var _len2 = arguments.length, b = new Array(_len2), _key2 = 0; _key2 < _len2; _key2++) {
      b[_key2] = arguments[_key2];
    }
    return hmacFn.apply(void 0, [k, v].concat(b));
  }; // hmac(k)(v, ...values)
  var reseed = function reseed() {
    var seed = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : u8n();
    // HMAC-DRBG reseed() function. Steps D-G
    k = h(u8fr([0x00]), seed); // k = hmac(k || v || 0x00 || seed)
    v = h(); // v = hmac(k || v)
    if (seed.length === 0) return;
    k = h(u8fr([0x01]), seed); // k = hmac(k || v || 0x01 || seed)
    v = h(); // v = hmac(k || v)
  };
  var gen = function gen() {
    // HMAC-DRBG generate() function
    if (i++ >= 1000) throw new Error('drbg: tried 1000 values');
    var len = 0;
    var out = [];
    while (len < qByteLen) {
      v = h();
      var sl = v.slice();
      out.push(sl);
      len += v.length;
    }
    return concatBytes.apply(void 0, out);
  };
  var genUntil = function genUntil(seed, pred) {
    reset();
    reseed(seed); // Steps D-G
    var res = undefined; // Step H: grind until k is in [1..n-1]
    while (!(res = pred(gen()))) reseed();
    reset();
    return res;
  };
  return genUntil;
}
// Validating curves and fields
var validatorFns = {
  bigint: function bigint(val) {
    return typeof val === 'bigint';
  },
  "function": function _function(val) {
    return typeof val === 'function';
  },
  "boolean": function boolean(val) {
    return typeof val === 'boolean';
  },
  string: function string(val) {
    return typeof val === 'string';
  },
  stringOrUint8Array: function stringOrUint8Array(val) {
    return typeof val === 'string' || val instanceof Uint8Array;
  },
  isSafeInteger: function isSafeInteger(val) {
    return Number.isSafeInteger(val);
  },
  array: function array(val) {
    return Array.isArray(val);
  },
  field: function field(val, object) {
    return object.Fp.isValid(val);
  },
  hash: function hash(val) {
    return typeof val === 'function' && Number.isSafeInteger(val.outputLen);
  }
};
// type Record<K extends string | number | symbol, T> = { [P in K]: T; }
function validateObject(object, validators) {
  var optValidators = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : {};
  var checkField = function checkField(fieldName, type, isOptional) {
    var checkVal = validatorFns[type];
    if (typeof checkVal !== 'function') throw new Error("Invalid validator \"".concat(type, "\", expected function"));
    var val = object[fieldName];
    if (isOptional && val === undefined) return;
    if (!checkVal(val, object)) {
      throw new Error("Invalid param ".concat(String(fieldName), "=").concat(val, " (").concat(_typeof(val), "), expected ").concat(type));
    }
  };
  for (var _i = 0, _Object$entries = Object.entries(validators); _i < _Object$entries.length; _i++) {
    var _Object$entries$_i = _slicedToArray(_Object$entries[_i], 2),
      fieldName = _Object$entries$_i[0],
      type = _Object$entries$_i[1];
    checkField(fieldName, type, false);
  }
  for (var _i2 = 0, _Object$entries2 = Object.entries(optValidators); _i2 < _Object$entries2.length; _i2++) {
    var _Object$entries2$_i = _slicedToArray(_Object$entries2[_i2], 2),
      _fieldName = _Object$entries2$_i[0],
      _type = _Object$entries2$_i[1];
    checkField(_fieldName, _type, true);
  }
  return object;
}
// validate type tests
// const o: { a: number; b: number; c: number } = { a: 1, b: 5, c: 6 };
// const z0 = validateObject(o, { a: 'isSafeInteger' }, { c: 'bigint' }); // Ok!
// // Should fail type-check
// const z1 = validateObject(o, { a: 'tmp' }, { c: 'zz' });
// const z2 = validateObject(o, { a: 'isSafeInteger' }, { c: 'zz' });
// const z3 = validateObject(o, { test: 'boolean', z: 'bug' });
// const z4 = validateObject(o, { a: 'boolean', z: 'bug' });

var ut = /*#__PURE__*/Object.freeze({
  __proto__: null,
  bitGet: bitGet,
  bitLen: bitLen,
  bitMask: bitMask,
  bitSet: bitSet,
  bytesToHex: bytesToHex,
  bytesToNumberBE: bytesToNumberBE,
  bytesToNumberLE: bytesToNumberLE,
  concatBytes: concatBytes,
  createHmacDrbg: createHmacDrbg,
  ensureBytes: ensureBytes,
  equalBytes: equalBytes,
  hexToBytes: hexToBytes,
  hexToNumber: hexToNumber,
  numberToBytesBE: numberToBytesBE,
  numberToBytesLE: numberToBytesLE,
  numberToHexUnpadded: numberToHexUnpadded,
  numberToVarBytesBE: numberToVarBytesBE,
  utf8ToBytes: utf8ToBytes,
  validateObject: validateObject
});

// prettier-ignore
var _0n$2 = BigInt(0),
  _1n$3 = BigInt(1),
  _2n$1 = BigInt(2),
  _3n$1 = BigInt(3);
// prettier-ignore
var _4n = BigInt(4),
  _5n = BigInt(5),
  _8n = BigInt(8);
// prettier-ignore
BigInt(9);
  BigInt(16);
// Calculates a modulo b
function mod(a, b) {
  var result = a % b;
  return result >= _0n$2 ? result : b + result;
}
/**
 * Efficiently raise num to power and do modular division.
 * Unsafe in some contexts: uses ladder, so can expose bigint bits.
 * @example
 * pow(2n, 6n, 11n) // 64n % 11n == 9n
 */
// TODO: use field version && remove
function pow(num, power, modulo) {
  if (modulo <= _0n$2 || power < _0n$2) throw new Error('Expected power/modulo > 0');
  if (modulo === _1n$3) return _0n$2;
  var res = _1n$3;
  while (power > _0n$2) {
    if (power & _1n$3) res = res * num % modulo;
    num = num * num % modulo;
    power >>= _1n$3;
  }
  return res;
}
// Does x ^ (2 ^ power) mod p. pow2(30, 4) == 30 ^ (2 ^ 4)
function pow2(x, power, modulo) {
  var res = x;
  while (power-- > _0n$2) {
    res *= res;
    res %= modulo;
  }
  return res;
}
// Inverses number over modulo
function invert(number, modulo) {
  if (number === _0n$2 || modulo <= _0n$2) {
    throw new Error("invert: expected positive integers, got n=".concat(number, " mod=").concat(modulo));
  }
  // Euclidean GCD https://brilliant.org/wiki/extended-euclidean-algorithm/
  // Fermat's little theorem "CT-like" version inv(n) = n^(m-2) mod m is 30x slower.
  var a = mod(number, modulo);
  var b = modulo;
  // prettier-ignore
  var x = _0n$2,
    u = _1n$3;
  while (a !== _0n$2) {
    // JIT applies optimization if those two lines follow each other
    var q = b / a;
    var r = b % a;
    var m = x - u * q;
    // prettier-ignore
    b = a, a = r, x = u, u = m;
  }
  var gcd = b;
  if (gcd !== _1n$3) throw new Error('invert: does not exist');
  return mod(x, modulo);
}
/**
 * Tonelli-Shanks square root search algorithm.
 * 1. https://eprint.iacr.org/2012/685.pdf (page 12)
 * 2. Square Roots from 1; 24, 51, 10 to Dan Shanks
 * Will start an infinite loop if field order P is not prime.
 * @param P field order
 * @returns function that takes field Fp (created from P) and number n
 */
function tonelliShanks(P) {
  // Legendre constant: used to calculate Legendre symbol (a | p),
  // which denotes the value of a^((p-1)/2) (mod p).
  // (a | p) ≡ 1    if a is a square (mod p)
  // (a | p) ≡ -1   if a is not a square (mod p)
  // (a | p) ≡ 0    if a ≡ 0 (mod p)
  var legendreC = (P - _1n$3) / _2n$1;
  var Q, S, Z;
  // Step 1: By factoring out powers of 2 from p - 1,
  // find q and s such that p - 1 = q*(2^s) with q odd
  for (Q = P - _1n$3, S = 0; Q % _2n$1 === _0n$2; Q /= _2n$1, S++);
  // Step 2: Select a non-square z such that (z | p) ≡ -1 and set c ≡ zq
  for (Z = _2n$1; Z < P && pow(Z, legendreC, P) !== P - _1n$3; Z++);
  // Fast-path
  if (S === 1) {
    var p1div4 = (P + _1n$3) / _4n;
    return function tonelliFast(Fp, n) {
      var root = Fp.pow(n, p1div4);
      if (!Fp.eql(Fp.sqr(root), n)) throw new Error('Cannot find square root');
      return root;
    };
  }
  // Slow-path
  var Q1div2 = (Q + _1n$3) / _2n$1;
  return function tonelliSlow(Fp, n) {
    // Step 0: Check that n is indeed a square: (n | p) should not be ≡ -1
    if (Fp.pow(n, legendreC) === Fp.neg(Fp.ONE)) throw new Error('Cannot find square root');
    var r = S;
    // TODO: will fail at Fp2/etc
    var g = Fp.pow(Fp.mul(Fp.ONE, Z), Q); // will update both x and b
    var x = Fp.pow(n, Q1div2); // first guess at the square root
    var b = Fp.pow(n, Q); // first guess at the fudge factor
    while (!Fp.eql(b, Fp.ONE)) {
      if (Fp.eql(b, Fp.ZERO)) return Fp.ZERO; // https://en.wikipedia.org/wiki/Tonelli%E2%80%93Shanks_algorithm (4. If t = 0, return r = 0)
      // Find m such b^(2^m)==1
      var m = 1;
      for (var t2 = Fp.sqr(b); m < r; m++) {
        if (Fp.eql(t2, Fp.ONE)) break;
        t2 = Fp.sqr(t2); // t2 *= t2
      }
      // NOTE: r-m-1 can be bigger than 32, need to convert to bigint before shift, otherwise there will be overflow
      var ge = Fp.pow(g, _1n$3 << BigInt(r - m - 1)); // ge = 2^(r-m-1)
      g = Fp.sqr(ge); // g = ge * ge
      x = Fp.mul(x, ge); // x *= ge
      b = Fp.mul(b, g); // b *= g
      r = m;
    }
    return x;
  };
}
function FpSqrt(P) {
  // NOTE: different algorithms can give different roots, it is up to user to decide which one they want.
  // For example there is FpSqrtOdd/FpSqrtEven to choice root based on oddness (used for hash-to-curve).
  // P ≡ 3 (mod 4)
  // √n = n^((P+1)/4)
  if (P % _4n === _3n$1) {
    // Not all roots possible!
    // const ORDER =
    //   0x1a0111ea397fe69a4b1ba7b6434bacd764774b84f38512bf6730d2a0f6b0f6241eabfffeb153ffffb9feffffffffaaabn;
    // const NUM = 72057594037927816n;
    var p1div4 = (P + _1n$3) / _4n;
    return function sqrt3mod4(Fp, n) {
      var root = Fp.pow(n, p1div4);
      // Throw if root**2 != n
      if (!Fp.eql(Fp.sqr(root), n)) throw new Error('Cannot find square root');
      return root;
    };
  }
  // Atkin algorithm for q ≡ 5 (mod 8), https://eprint.iacr.org/2012/685.pdf (page 10)
  if (P % _8n === _5n) {
    var c1 = (P - _5n) / _8n;
    return function sqrt5mod8(Fp, n) {
      var n2 = Fp.mul(n, _2n$1);
      var v = Fp.pow(n2, c1);
      var nv = Fp.mul(n, v);
      var i = Fp.mul(Fp.mul(nv, _2n$1), v);
      var root = Fp.mul(nv, Fp.sub(i, Fp.ONE));
      if (!Fp.eql(Fp.sqr(root), n)) throw new Error('Cannot find square root');
      return root;
    };
  }
  // Other cases: Tonelli-Shanks algorithm
  return tonelliShanks(P);
}
// prettier-ignore
var FIELD_FIELDS = ['create', 'isValid', 'is0', 'neg', 'inv', 'sqrt', 'sqr', 'eql', 'add', 'sub', 'mul', 'pow', 'div', 'addN', 'subN', 'mulN', 'sqrN'];
function validateField(field) {
  var initial = {
    ORDER: 'bigint',
    MASK: 'bigint',
    BYTES: 'isSafeInteger',
    BITS: 'isSafeInteger'
  };
  var opts = FIELD_FIELDS.reduce(function (map, val) {
    map[val] = 'function';
    return map;
  }, initial);
  return validateObject(field, opts);
}
// Generic field functions
/**
 * Same as `pow` but for Fp: non-constant-time.
 * Unsafe in some contexts: uses ladder, so can expose bigint bits.
 */
function FpPow(f, num, power) {
  // Should have same speed as pow for bigints
  // TODO: benchmark!
  if (power < _0n$2) throw new Error('Expected power > 0');
  if (power === _0n$2) return f.ONE;
  if (power === _1n$3) return num;
  var p = f.ONE;
  var d = num;
  while (power > _0n$2) {
    if (power & _1n$3) p = f.mul(p, d);
    d = f.sqr(d);
    power >>= _1n$3;
  }
  return p;
}
/**
 * Efficiently invert an array of Field elements.
 * `inv(0)` will return `undefined` here: make sure to throw an error.
 */
function FpInvertBatch(f, nums) {
  var tmp = new Array(nums.length);
  // Walk from first to last, multiply them by each other MOD p
  var lastMultiplied = nums.reduce(function (acc, num, i) {
    if (f.is0(num)) return acc;
    tmp[i] = acc;
    return f.mul(acc, num);
  }, f.ONE);
  // Invert last element
  var inverted = f.inv(lastMultiplied);
  // Walk from last to first, multiply them by inverted each other MOD p
  nums.reduceRight(function (acc, num, i) {
    if (f.is0(num)) return acc;
    tmp[i] = f.mul(acc, tmp[i]);
    return f.mul(acc, num);
  }, inverted);
  return tmp;
}
// CURVE.n lengths
function nLength(n, nBitLength) {
  // Bit size, byte size of CURVE.n
  var _nBitLength = nBitLength !== undefined ? nBitLength : n.toString(2).length;
  var nByteLength = Math.ceil(_nBitLength / 8);
  return {
    nBitLength: _nBitLength,
    nByteLength: nByteLength
  };
}
/**
 * Initializes a finite field over prime. **Non-primes are not supported.**
 * Do not init in loop: slow. Very fragile: always run a benchmark on a change.
 * Major performance optimizations:
 * * a) denormalized operations like mulN instead of mul
 * * b) same object shape: never add or remove keys
 * * c) Object.freeze
 * @param ORDER prime positive bigint
 * @param bitLen how many bits the field consumes
 * @param isLE (def: false) if encoding / decoding should be in little-endian
 * @param redef optional faster redefinitions of sqrt and other methods
 */
function Field(ORDER, bitLen) {
  var isLE = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : false;
  var redef = arguments.length > 3 && arguments[3] !== undefined ? arguments[3] : {};
  if (ORDER <= _0n$2) throw new Error("Expected Field ORDER > 0, got ".concat(ORDER));
  var _nLength = nLength(ORDER, bitLen),
    BITS = _nLength.nBitLength,
    BYTES = _nLength.nByteLength;
  if (BYTES > 2048) throw new Error('Field lengths over 2048 bytes are not supported');
  var sqrtP = FpSqrt(ORDER);
  var f = Object.freeze({
    ORDER: ORDER,
    BITS: BITS,
    BYTES: BYTES,
    MASK: bitMask(BITS),
    ZERO: _0n$2,
    ONE: _1n$3,
    create: function create(num) {
      return mod(num, ORDER);
    },
    isValid: function isValid(num) {
      if (typeof num !== 'bigint') throw new Error("Invalid field element: expected bigint, got ".concat(_typeof(num)));
      return _0n$2 <= num && num < ORDER; // 0 is valid element, but it's not invertible
    },
    is0: function is0(num) {
      return num === _0n$2;
    },
    isOdd: function isOdd(num) {
      return (num & _1n$3) === _1n$3;
    },
    neg: function neg(num) {
      return mod(-num, ORDER);
    },
    eql: function eql(lhs, rhs) {
      return lhs === rhs;
    },
    sqr: function sqr(num) {
      return mod(num * num, ORDER);
    },
    add: function add(lhs, rhs) {
      return mod(lhs + rhs, ORDER);
    },
    sub: function sub(lhs, rhs) {
      return mod(lhs - rhs, ORDER);
    },
    mul: function mul(lhs, rhs) {
      return mod(lhs * rhs, ORDER);
    },
    pow: function pow(num, power) {
      return FpPow(f, num, power);
    },
    div: function div(lhs, rhs) {
      return mod(lhs * invert(rhs, ORDER), ORDER);
    },
    // Same as above, but doesn't normalize
    sqrN: function sqrN(num) {
      return num * num;
    },
    addN: function addN(lhs, rhs) {
      return lhs + rhs;
    },
    subN: function subN(lhs, rhs) {
      return lhs - rhs;
    },
    mulN: function mulN(lhs, rhs) {
      return lhs * rhs;
    },
    inv: function inv(num) {
      return invert(num, ORDER);
    },
    sqrt: redef.sqrt || function (n) {
      return sqrtP(f, n);
    },
    invertBatch: function invertBatch(lst) {
      return FpInvertBatch(f, lst);
    },
    // TODO: do we really need constant cmov?
    // We don't have const-time bigints anyway, so probably will be not very useful
    cmov: function cmov(a, b, c) {
      return c ? b : a;
    },
    toBytes: function toBytes(num) {
      return isLE ? numberToBytesLE(num, BYTES) : numberToBytesBE(num, BYTES);
    },
    fromBytes: function fromBytes(bytes) {
      if (bytes.length !== BYTES) throw new Error("Fp.fromBytes: expected ".concat(BYTES, ", got ").concat(bytes.length));
      return isLE ? bytesToNumberLE(bytes) : bytesToNumberBE(bytes);
    }
  });
  return Object.freeze(f);
}
/**
 * Returns total number of bytes consumed by the field element.
 * For example, 32 bytes for usual 256-bit weierstrass curve.
 * @param fieldOrder number of field elements, usually CURVE.n
 * @returns byte length of field
 */
function getFieldBytesLength(fieldOrder) {
  if (typeof fieldOrder !== 'bigint') throw new Error('field order must be bigint');
  var bitLength = fieldOrder.toString(2).length;
  return Math.ceil(bitLength / 8);
}
/**
 * Returns minimal amount of bytes that can be safely reduced
 * by field order.
 * Should be 2^-128 for 128-bit curve such as P256.
 * @param fieldOrder number of field elements, usually CURVE.n
 * @returns byte length of target hash
 */
function getMinHashLength(fieldOrder) {
  var length = getFieldBytesLength(fieldOrder);
  return length + Math.ceil(length / 2);
}
/**
 * "Constant-time" private key generation utility.
 * Can take (n + n/2) or more bytes of uniform input e.g. from CSPRNG or KDF
 * and convert them into private scalar, with the modulo bias being negligible.
 * Needs at least 48 bytes of input for 32-byte private key.
 * https://research.kudelskisecurity.com/2020/07/28/the-definitive-guide-to-modulo-bias-and-how-to-avoid-it/
 * FIPS 186-5, A.2 https://csrc.nist.gov/publications/detail/fips/186/5/final
 * RFC 9380, https://www.rfc-editor.org/rfc/rfc9380#section-5
 * @param hash hash output from SHA3 or a similar function
 * @param groupOrder size of subgroup - (e.g. secp256k1.CURVE.n)
 * @param isLE interpret hash bytes as LE num
 * @returns valid private scalar
 */
function mapHashToField(key, fieldOrder) {
  var isLE = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : false;
  var len = key.length;
  var fieldLen = getFieldBytesLength(fieldOrder);
  var minLen = getMinHashLength(fieldOrder);
  // No small numbers: need to understand bias story. No huge numbers: easier to detect JS timings.
  if (len < 16 || len < minLen || len > 1024) throw new Error("expected ".concat(minLen, "-1024 bytes of input, got ").concat(len));
  var num = isLE ? bytesToNumberBE(key) : bytesToNumberLE(key);
  // `mod(x, 11)` can sometimes produce 0. `mod(x, 10) + 1` is the same, but no 0
  var reduced = mod(num, fieldOrder - _1n$3) + _1n$3;
  return isLE ? numberToBytesLE(reduced, fieldLen) : numberToBytesBE(reduced, fieldLen);
}

var _0n$1 = BigInt(0);
var _1n$2 = BigInt(1);
// Elliptic curve multiplication of Point by scalar. Fragile.
// Scalars should always be less than curve order: this should be checked inside of a curve itself.
// Creates precomputation tables for fast multiplication:
// - private scalar is split by fixed size windows of W bits
// - every window point is collected from window's table & added to accumulator
// - since windows are different, same point inside tables won't be accessed more than once per calc
// - each multiplication is 'Math.ceil(CURVE_ORDER / 𝑊) + 1' point additions (fixed for any scalar)
// - +1 window is neccessary for wNAF
// - wNAF reduces table size: 2x less memory + 2x faster generation, but 10% slower multiplication
// TODO: Research returning 2d JS array of windows, instead of a single window. This would allow
// windows to be in different memory locations
function wNAF(c, bits) {
  var constTimeNegate = function constTimeNegate(condition, item) {
    var neg = item.negate();
    return condition ? neg : item;
  };
  var opts = function opts(W) {
    var windows = Math.ceil(bits / W) + 1; // +1, because
    var windowSize = Math.pow(2, W - 1); // -1 because we skip zero
    return {
      windows: windows,
      windowSize: windowSize
    };
  };
  return {
    constTimeNegate: constTimeNegate,
    // non-const time multiplication ladder
    unsafeLadder: function unsafeLadder(elm, n) {
      var p = c.ZERO;
      var d = elm;
      while (n > _0n$1) {
        if (n & _1n$2) p = p.add(d);
        d = d["double"]();
        n >>= _1n$2;
      }
      return p;
    },
    /**
     * Creates a wNAF precomputation window. Used for caching.
     * Default window size is set by `utils.precompute()` and is equal to 8.
     * Number of precomputed points depends on the curve size:
     * 2^(𝑊−1) * (Math.ceil(𝑛 / 𝑊) + 1), where:
     * - 𝑊 is the window size
     * - 𝑛 is the bitlength of the curve order.
     * For a 256-bit curve and window size 8, the number of precomputed points is 128 * 33 = 4224.
     * @returns precomputed point tables flattened to a single array
     */
    precomputeWindow: function precomputeWindow(elm, W) {
      var _opts = opts(W),
        windows = _opts.windows,
        windowSize = _opts.windowSize;
      var points = [];
      var p = elm;
      var base = p;
      for (var window = 0; window < windows; window++) {
        base = p;
        points.push(base);
        // =1, because we skip zero
        for (var i = 1; i < windowSize; i++) {
          base = base.add(p);
          points.push(base);
        }
        p = base["double"]();
      }
      return points;
    },
    /**
     * Implements ec multiplication using precomputed tables and w-ary non-adjacent form.
     * @param W window size
     * @param precomputes precomputed tables
     * @param n scalar (we don't check here, but should be less than curve order)
     * @returns real and fake (for const-time) points
     */
    wNAF: function wNAF(W, precomputes, n) {
      // TODO: maybe check that scalar is less than group order? wNAF behavious is undefined otherwise
      // But need to carefully remove other checks before wNAF. ORDER == bits here
      var _opts2 = opts(W),
        windows = _opts2.windows,
        windowSize = _opts2.windowSize;
      var p = c.ZERO;
      var f = c.BASE;
      var mask = BigInt(Math.pow(2, W) - 1); // Create mask with W ones: 0b1111 for W=4 etc.
      var maxNumber = Math.pow(2, W);
      var shiftBy = BigInt(W);
      for (var window = 0; window < windows; window++) {
        var offset = window * windowSize;
        // Extract W bits.
        var wbits = Number(n & mask);
        // Shift number by W bits.
        n >>= shiftBy;
        // If the bits are bigger than max size, we'll split those.
        // +224 => 256 - 32
        if (wbits > windowSize) {
          wbits -= maxNumber;
          n += _1n$2;
        }
        // This code was first written with assumption that 'f' and 'p' will never be infinity point:
        // since each addition is multiplied by 2 ** W, it cannot cancel each other. However,
        // there is negate now: it is possible that negated element from low value
        // would be the same as high element, which will create carry into next window.
        // It's not obvious how this can fail, but still worth investigating later.
        // Check if we're onto Zero point.
        // Add random point inside current window to f.
        var offset1 = offset;
        var offset2 = offset + Math.abs(wbits) - 1; // -1 because we skip zero
        var cond1 = window % 2 !== 0;
        var cond2 = wbits < 0;
        if (wbits === 0) {
          // The most important part for const-time getPublicKey
          f = f.add(constTimeNegate(cond1, precomputes[offset1]));
        } else {
          p = p.add(constTimeNegate(cond2, precomputes[offset2]));
        }
      }
      // JIT-compiler should not eliminate f here, since it will later be used in normalizeZ()
      // Even if the variable is still unused, there are some checks which will
      // throw an exception, so compiler needs to prove they won't happen, which is hard.
      // At this point there is a way to F be infinity-point even if p is not,
      // which makes it less const-time: around 1 bigint multiply.
      return {
        p: p,
        f: f
      };
    },
    wNAFCached: function wNAFCached(P, precomputesMap, n, transform) {
      // @ts-ignore
      var W = P._WINDOW_SIZE || 1;
      // Calculate precomputes on a first run, reuse them after
      var comp = precomputesMap.get(P);
      if (!comp) {
        comp = this.precomputeWindow(P, W);
        if (W !== 1) {
          precomputesMap.set(P, transform(comp));
        }
      }
      return this.wNAF(W, comp, n);
    }
  };
}
function validateBasic(curve) {
  validateField(curve.Fp);
  validateObject(curve, {
    n: 'bigint',
    h: 'bigint',
    Gx: 'field',
    Gy: 'field'
  }, {
    nBitLength: 'isSafeInteger',
    nByteLength: 'isSafeInteger'
  });
  // Set defaults
  return Object.freeze(_objectSpread2(_objectSpread2(_objectSpread2({}, nLength(curve.n, curve.nBitLength)), curve), {
    p: curve.Fp.ORDER
  }));
}

function validatePointOpts(curve) {
  var opts = validateBasic(curve);
  validateObject(opts, {
    a: 'field',
    b: 'field'
  }, {
    allowedPrivateKeyLengths: 'array',
    wrapPrivateKey: 'boolean',
    isTorsionFree: 'function',
    clearCofactor: 'function',
    allowInfinityPoint: 'boolean',
    fromBytes: 'function',
    toBytes: 'function'
  });
  var endo = opts.endo,
    Fp = opts.Fp,
    a = opts.a;
  if (endo) {
    if (!Fp.eql(a, Fp.ZERO)) {
      throw new Error('Endomorphism can only be defined for Koblitz curves that have a=0');
    }
    if (_typeof(endo) !== 'object' || typeof endo.beta !== 'bigint' || typeof endo.splitScalar !== 'function') {
      throw new Error('Expected endomorphism with beta: bigint and splitScalar: function');
    }
  }
  return Object.freeze(_objectSpread2({}, opts));
}
// ASN.1 DER encoding utilities
var b2n = bytesToNumberBE,
  h2b = hexToBytes;
var DER = {
  // asn.1 DER encoding utils
  Err: /*#__PURE__*/function (_Error) {
    function DERErr() {
      var m = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : '';
      _classCallCheck(this, DERErr);
      return _callSuper(this, DERErr, [m]);
    }
    _inherits(DERErr, _Error);
    return _createClass(DERErr);
  }(/*#__PURE__*/_wrapNativeSuper(Error)),
  _parseInt: function _parseInt(data) {
    var E = DER.Err;
    if (data.length < 2 || data[0] !== 0x02) throw new E('Invalid signature integer tag');
    var len = data[1];
    var res = data.subarray(2, len + 2);
    if (!len || res.length !== len) throw new E('Invalid signature integer: wrong length');
    // https://crypto.stackexchange.com/a/57734 Leftmost bit of first byte is 'negative' flag,
    // since we always use positive integers here. It must always be empty:
    // - add zero byte if exists
    // - if next byte doesn't have a flag, leading zero is not allowed (minimal encoding)
    if (res[0] & 128) throw new E('Invalid signature integer: negative');
    if (res[0] === 0x00 && !(res[1] & 128)) throw new E('Invalid signature integer: unnecessary leading zero');
    return {
      d: b2n(res),
      l: data.subarray(len + 2)
    }; // d is data, l is left
  },
  toSig: function toSig(hex) {
    // parse DER signature
    var E = DER.Err;
    var data = typeof hex === 'string' ? h2b(hex) : hex;
    if (!(data instanceof Uint8Array)) throw new Error('ui8a expected');
    var l = data.length;
    if (l < 2 || data[0] != 0x30) throw new E('Invalid signature tag');
    if (data[1] !== l - 2) throw new E('Invalid signature: incorrect length');
    var _DER$_parseInt = DER._parseInt(data.subarray(2)),
      r = _DER$_parseInt.d,
      sBytes = _DER$_parseInt.l;
    var _DER$_parseInt2 = DER._parseInt(sBytes),
      s = _DER$_parseInt2.d,
      rBytesLeft = _DER$_parseInt2.l;
    if (rBytesLeft.length) throw new E('Invalid signature: left bytes after parsing');
    return {
      r: r,
      s: s
    };
  },
  hexFromSig: function hexFromSig(sig) {
    // Add leading zero if first byte has negative bit enabled. More details in '_parseInt'
    var slice = function slice(s) {
      return Number.parseInt(s[0], 16) & 8 ? '00' + s : s;
    };
    var h = function h(num) {
      var hex = num.toString(16);
      return hex.length & 1 ? "0".concat(hex) : hex;
    };
    var s = slice(h(sig.s));
    var r = slice(h(sig.r));
    var shl = s.length / 2;
    var rhl = r.length / 2;
    var sl = h(shl);
    var rl = h(rhl);
    return "30".concat(h(rhl + shl + 4), "02").concat(rl).concat(r, "02").concat(sl).concat(s);
  }
};
// Be friendly to bad ECMAScript parsers by not using bigint literals
// prettier-ignore
var _0n = BigInt(0),
  _1n$1 = BigInt(1);
  BigInt(2);
  var _3n = BigInt(3);
  BigInt(4);
function weierstrassPoints(opts) {
  var CURVE = validatePointOpts(opts);
  var Fp = CURVE.Fp; // All curves has same field / group length as for now, but they can differ
  var toBytes = CURVE.toBytes || function (_c, point, _isCompressed) {
    var a = point.toAffine();
    return concatBytes(Uint8Array.from([0x04]), Fp.toBytes(a.x), Fp.toBytes(a.y));
  };
  var fromBytes = CURVE.fromBytes || function (bytes) {
    // const head = bytes[0];
    var tail = bytes.subarray(1);
    // if (head !== 0x04) throw new Error('Only non-compressed encoding is supported');
    var x = Fp.fromBytes(tail.subarray(0, Fp.BYTES));
    var y = Fp.fromBytes(tail.subarray(Fp.BYTES, 2 * Fp.BYTES));
    return {
      x: x,
      y: y
    };
  };
  /**
   * y² = x³ + ax + b: Short weierstrass curve formula
   * @returns y²
   */
  function weierstrassEquation(x) {
    var a = CURVE.a,
      b = CURVE.b;
    var x2 = Fp.sqr(x); // x * x
    var x3 = Fp.mul(x2, x); // x2 * x
    return Fp.add(Fp.add(x3, Fp.mul(x, a)), b); // x3 + a * x + b
  }
  // Validate whether the passed curve params are valid.
  // We check if curve equation works for generator point.
  // `assertValidity()` won't work: `isTorsionFree()` is not available at this point in bls12-381.
  // ProjectivePoint class has not been initialized yet.
  if (!Fp.eql(Fp.sqr(CURVE.Gy), weierstrassEquation(CURVE.Gx))) throw new Error('bad generator point: equation left != right');
  // Valid group elements reside in range 1..n-1
  function isWithinCurveOrder(num) {
    return typeof num === 'bigint' && _0n < num && num < CURVE.n;
  }
  function assertGE(num) {
    if (!isWithinCurveOrder(num)) throw new Error('Expected valid bigint: 0 < bigint < curve.n');
  }
  // Validates if priv key is valid and converts it to bigint.
  // Supports options allowedPrivateKeyLengths and wrapPrivateKey.
  function normPrivateKeyToScalar(key) {
    var lengths = CURVE.allowedPrivateKeyLengths,
      nByteLength = CURVE.nByteLength,
      wrapPrivateKey = CURVE.wrapPrivateKey,
      n = CURVE.n;
    if (lengths && typeof key !== 'bigint') {
      if (key instanceof Uint8Array) key = bytesToHex(key);
      // Normalize to hex string, pad. E.g. P521 would norm 130-132 char hex to 132-char bytes
      if (typeof key !== 'string' || !lengths.includes(key.length)) throw new Error('Invalid key');
      key = key.padStart(nByteLength * 2, '0');
    }
    var num;
    try {
      num = typeof key === 'bigint' ? key : bytesToNumberBE(ensureBytes('private key', key, nByteLength));
    } catch (error) {
      throw new Error("private key must be ".concat(nByteLength, " bytes, hex or bigint, not ").concat(_typeof(key)));
    }
    if (wrapPrivateKey) num = mod(num, n); // disabled by default, enabled for BLS
    assertGE(num); // num in range [1..N-1]
    return num;
  }
  var pointPrecomputes = new Map();
  function assertPrjPoint(other) {
    if (!(other instanceof Point)) throw new Error('ProjectivePoint expected');
  }
  /**
   * Projective Point works in 3d / projective (homogeneous) coordinates: (x, y, z) ∋ (x=x/z, y=y/z)
   * Default Point works in 2d / affine coordinates: (x, y)
   * We're doing calculations in projective, because its operations don't require costly inversion.
   */
  var Point = /*#__PURE__*/function () {
    function Point(px, py, pz) {
      _classCallCheck(this, Point);
      this.px = px;
      this.py = py;
      this.pz = pz;
      if (px == null || !Fp.isValid(px)) throw new Error('x required');
      if (py == null || !Fp.isValid(py)) throw new Error('y required');
      if (pz == null || !Fp.isValid(pz)) throw new Error('z required');
    }
    // Does not validate if the point is on-curve.
    // Use fromHex instead, or call assertValidity() later.
    return _createClass(Point, [{
      key: "x",
      get: function get() {
        return this.toAffine().x;
      }
    }, {
      key: "y",
      get: function get() {
        return this.toAffine().y;
      }
      /**
       * Takes a bunch of Projective Points but executes only one
       * inversion on all of them. Inversion is very slow operation,
       * so this improves performance massively.
       * Optimization: converts a list of projective points to a list of identical points with Z=1.
       */
    }, {
      key: "_setWindowSize",
      value:
      // "Private method", don't use it directly
      function _setWindowSize(windowSize) {
        this._WINDOW_SIZE = windowSize;
        pointPrecomputes["delete"](this);
      }
      // A point on curve is valid if it conforms to equation.
    }, {
      key: "assertValidity",
      value: function assertValidity() {
        if (this.is0()) {
          // (0, 1, 0) aka ZERO is invalid in most contexts.
          // In BLS, ZERO can be serialized, so we allow it.
          // (0, 0, 0) is wrong representation of ZERO and is always invalid.
          if (CURVE.allowInfinityPoint && !Fp.is0(this.py)) return;
          throw new Error('bad point: ZERO');
        }
        // Some 3rd-party test vectors require different wording between here & `fromCompressedHex`
        var _this$toAffine = this.toAffine(),
          x = _this$toAffine.x,
          y = _this$toAffine.y;
        // Check if x, y are valid field elements
        if (!Fp.isValid(x) || !Fp.isValid(y)) throw new Error('bad point: x or y not FE');
        var left = Fp.sqr(y); // y²
        var right = weierstrassEquation(x); // x³ + ax + b
        if (!Fp.eql(left, right)) throw new Error('bad point: equation left != right');
        if (!this.isTorsionFree()) throw new Error('bad point: not in prime-order subgroup');
      }
    }, {
      key: "hasEvenY",
      value: function hasEvenY() {
        var _this$toAffine2 = this.toAffine(),
          y = _this$toAffine2.y;
        if (Fp.isOdd) return !Fp.isOdd(y);
        throw new Error("Field doesn't support isOdd");
      }
      /**
       * Compare one point to another.
       */
    }, {
      key: "equals",
      value: function equals(other) {
        assertPrjPoint(other);
        var X1 = this.px,
          Y1 = this.py,
          Z1 = this.pz;
        var X2 = other.px,
          Y2 = other.py,
          Z2 = other.pz;
        var U1 = Fp.eql(Fp.mul(X1, Z2), Fp.mul(X2, Z1));
        var U2 = Fp.eql(Fp.mul(Y1, Z2), Fp.mul(Y2, Z1));
        return U1 && U2;
      }
      /**
       * Flips point to one corresponding to (x, -y) in Affine coordinates.
       */
    }, {
      key: "negate",
      value: function negate() {
        return new Point(this.px, Fp.neg(this.py), this.pz);
      }
      // Renes-Costello-Batina exception-free doubling formula.
      // There is 30% faster Jacobian formula, but it is not complete.
      // https://eprint.iacr.org/2015/1060, algorithm 3
      // Cost: 8M + 3S + 3*a + 2*b3 + 15add.
    }, {
      key: "double",
      value: function _double() {
        var a = CURVE.a,
          b = CURVE.b;
        var b3 = Fp.mul(b, _3n);
        var X1 = this.px,
          Y1 = this.py,
          Z1 = this.pz;
        var X3 = Fp.ZERO,
          Y3 = Fp.ZERO,
          Z3 = Fp.ZERO; // prettier-ignore
        var t0 = Fp.mul(X1, X1); // step 1
        var t1 = Fp.mul(Y1, Y1);
        var t2 = Fp.mul(Z1, Z1);
        var t3 = Fp.mul(X1, Y1);
        t3 = Fp.add(t3, t3); // step 5
        Z3 = Fp.mul(X1, Z1);
        Z3 = Fp.add(Z3, Z3);
        X3 = Fp.mul(a, Z3);
        Y3 = Fp.mul(b3, t2);
        Y3 = Fp.add(X3, Y3); // step 10
        X3 = Fp.sub(t1, Y3);
        Y3 = Fp.add(t1, Y3);
        Y3 = Fp.mul(X3, Y3);
        X3 = Fp.mul(t3, X3);
        Z3 = Fp.mul(b3, Z3); // step 15
        t2 = Fp.mul(a, t2);
        t3 = Fp.sub(t0, t2);
        t3 = Fp.mul(a, t3);
        t3 = Fp.add(t3, Z3);
        Z3 = Fp.add(t0, t0); // step 20
        t0 = Fp.add(Z3, t0);
        t0 = Fp.add(t0, t2);
        t0 = Fp.mul(t0, t3);
        Y3 = Fp.add(Y3, t0);
        t2 = Fp.mul(Y1, Z1); // step 25
        t2 = Fp.add(t2, t2);
        t0 = Fp.mul(t2, t3);
        X3 = Fp.sub(X3, t0);
        Z3 = Fp.mul(t2, t1);
        Z3 = Fp.add(Z3, Z3); // step 30
        Z3 = Fp.add(Z3, Z3);
        return new Point(X3, Y3, Z3);
      }
      // Renes-Costello-Batina exception-free addition formula.
      // There is 30% faster Jacobian formula, but it is not complete.
      // https://eprint.iacr.org/2015/1060, algorithm 1
      // Cost: 12M + 0S + 3*a + 3*b3 + 23add.
    }, {
      key: "add",
      value: function add(other) {
        assertPrjPoint(other);
        var X1 = this.px,
          Y1 = this.py,
          Z1 = this.pz;
        var X2 = other.px,
          Y2 = other.py,
          Z2 = other.pz;
        var X3 = Fp.ZERO,
          Y3 = Fp.ZERO,
          Z3 = Fp.ZERO; // prettier-ignore
        var a = CURVE.a;
        var b3 = Fp.mul(CURVE.b, _3n);
        var t0 = Fp.mul(X1, X2); // step 1
        var t1 = Fp.mul(Y1, Y2);
        var t2 = Fp.mul(Z1, Z2);
        var t3 = Fp.add(X1, Y1);
        var t4 = Fp.add(X2, Y2); // step 5
        t3 = Fp.mul(t3, t4);
        t4 = Fp.add(t0, t1);
        t3 = Fp.sub(t3, t4);
        t4 = Fp.add(X1, Z1);
        var t5 = Fp.add(X2, Z2); // step 10
        t4 = Fp.mul(t4, t5);
        t5 = Fp.add(t0, t2);
        t4 = Fp.sub(t4, t5);
        t5 = Fp.add(Y1, Z1);
        X3 = Fp.add(Y2, Z2); // step 15
        t5 = Fp.mul(t5, X3);
        X3 = Fp.add(t1, t2);
        t5 = Fp.sub(t5, X3);
        Z3 = Fp.mul(a, t4);
        X3 = Fp.mul(b3, t2); // step 20
        Z3 = Fp.add(X3, Z3);
        X3 = Fp.sub(t1, Z3);
        Z3 = Fp.add(t1, Z3);
        Y3 = Fp.mul(X3, Z3);
        t1 = Fp.add(t0, t0); // step 25
        t1 = Fp.add(t1, t0);
        t2 = Fp.mul(a, t2);
        t4 = Fp.mul(b3, t4);
        t1 = Fp.add(t1, t2);
        t2 = Fp.sub(t0, t2); // step 30
        t2 = Fp.mul(a, t2);
        t4 = Fp.add(t4, t2);
        t0 = Fp.mul(t1, t4);
        Y3 = Fp.add(Y3, t0);
        t0 = Fp.mul(t5, t4); // step 35
        X3 = Fp.mul(t3, X3);
        X3 = Fp.sub(X3, t0);
        t0 = Fp.mul(t3, t1);
        Z3 = Fp.mul(t5, Z3);
        Z3 = Fp.add(Z3, t0); // step 40
        return new Point(X3, Y3, Z3);
      }
    }, {
      key: "subtract",
      value: function subtract(other) {
        return this.add(other.negate());
      }
    }, {
      key: "is0",
      value: function is0() {
        return this.equals(Point.ZERO);
      }
    }, {
      key: "wNAF",
      value: function wNAF(n) {
        return wnaf.wNAFCached(this, pointPrecomputes, n, function (comp) {
          var toInv = Fp.invertBatch(comp.map(function (p) {
            return p.pz;
          }));
          return comp.map(function (p, i) {
            return p.toAffine(toInv[i]);
          }).map(Point.fromAffine);
        });
      }
      /**
       * Non-constant-time multiplication. Uses double-and-add algorithm.
       * It's faster, but should only be used when you don't care about
       * an exposed private key e.g. sig verification, which works over *public* keys.
       */
    }, {
      key: "multiplyUnsafe",
      value: function multiplyUnsafe(n) {
        var I = Point.ZERO;
        if (n === _0n) return I;
        assertGE(n); // Will throw on 0
        if (n === _1n$1) return this;
        var endo = CURVE.endo;
        if (!endo) return wnaf.unsafeLadder(this, n);
        // Apply endomorphism
        var _endo$splitScalar = endo.splitScalar(n),
          k1neg = _endo$splitScalar.k1neg,
          k1 = _endo$splitScalar.k1,
          k2neg = _endo$splitScalar.k2neg,
          k2 = _endo$splitScalar.k2;
        var k1p = I;
        var k2p = I;
        var d = this;
        while (k1 > _0n || k2 > _0n) {
          if (k1 & _1n$1) k1p = k1p.add(d);
          if (k2 & _1n$1) k2p = k2p.add(d);
          d = d["double"]();
          k1 >>= _1n$1;
          k2 >>= _1n$1;
        }
        if (k1neg) k1p = k1p.negate();
        if (k2neg) k2p = k2p.negate();
        k2p = new Point(Fp.mul(k2p.px, endo.beta), k2p.py, k2p.pz);
        return k1p.add(k2p);
      }
      /**
       * Constant time multiplication.
       * Uses wNAF method. Windowed method may be 10% faster,
       * but takes 2x longer to generate and consumes 2x memory.
       * Uses precomputes when available.
       * Uses endomorphism for Koblitz curves.
       * @param scalar by which the point would be multiplied
       * @returns New point
       */
    }, {
      key: "multiply",
      value: function multiply(scalar) {
        assertGE(scalar);
        var n = scalar;
        var point, fake; // Fake point is used to const-time mult
        var endo = CURVE.endo;
        if (endo) {
          var _endo$splitScalar2 = endo.splitScalar(n),
            k1neg = _endo$splitScalar2.k1neg,
            k1 = _endo$splitScalar2.k1,
            k2neg = _endo$splitScalar2.k2neg,
            k2 = _endo$splitScalar2.k2;
          var _this$wNAF = this.wNAF(k1),
            k1p = _this$wNAF.p,
            f1p = _this$wNAF.f;
          var _this$wNAF2 = this.wNAF(k2),
            k2p = _this$wNAF2.p,
            f2p = _this$wNAF2.f;
          k1p = wnaf.constTimeNegate(k1neg, k1p);
          k2p = wnaf.constTimeNegate(k2neg, k2p);
          k2p = new Point(Fp.mul(k2p.px, endo.beta), k2p.py, k2p.pz);
          point = k1p.add(k2p);
          fake = f1p.add(f2p);
        } else {
          var _this$wNAF3 = this.wNAF(n),
            p = _this$wNAF3.p,
            f = _this$wNAF3.f;
          point = p;
          fake = f;
        }
        // Normalize `z` for both points, but return only real one
        return Point.normalizeZ([point, fake])[0];
      }
      /**
       * Efficiently calculate `aP + bQ`. Unsafe, can expose private key, if used incorrectly.
       * Not using Strauss-Shamir trick: precomputation tables are faster.
       * The trick could be useful if both P and Q are not G (not in our case).
       * @returns non-zero affine point
       */
    }, {
      key: "multiplyAndAddUnsafe",
      value: function multiplyAndAddUnsafe(Q, a, b) {
        var G = Point.BASE; // No Strauss-Shamir trick: we have 10% faster G precomputes
        var mul = function mul(P, a // Select faster multiply() method
        ) {
          return a === _0n || a === _1n$1 || !P.equals(G) ? P.multiplyUnsafe(a) : P.multiply(a);
        };
        var sum = mul(this, a).add(mul(Q, b));
        return sum.is0() ? undefined : sum;
      }
      // Converts Projective point to affine (x, y) coordinates.
      // Can accept precomputed Z^-1 - for example, from invertBatch.
      // (x, y, z) ∋ (x=x/z, y=y/z)
    }, {
      key: "toAffine",
      value: function toAffine(iz) {
        var x = this.px,
          y = this.py,
          z = this.pz;
        var is0 = this.is0();
        // If invZ was 0, we return zero point. However we still want to execute
        // all operations, so we replace invZ with a random number, 1.
        if (iz == null) iz = is0 ? Fp.ONE : Fp.inv(z);
        var ax = Fp.mul(x, iz);
        var ay = Fp.mul(y, iz);
        var zz = Fp.mul(z, iz);
        if (is0) return {
          x: Fp.ZERO,
          y: Fp.ZERO
        };
        if (!Fp.eql(zz, Fp.ONE)) throw new Error('invZ was invalid');
        return {
          x: ax,
          y: ay
        };
      }
    }, {
      key: "isTorsionFree",
      value: function isTorsionFree() {
        var cofactor = CURVE.h,
          isTorsionFree = CURVE.isTorsionFree;
        if (cofactor === _1n$1) return true; // No subgroups, always torsion-free
        if (isTorsionFree) return isTorsionFree(Point, this);
        throw new Error('isTorsionFree() has not been declared for the elliptic curve');
      }
    }, {
      key: "clearCofactor",
      value: function clearCofactor() {
        var cofactor = CURVE.h,
          clearCofactor = CURVE.clearCofactor;
        if (cofactor === _1n$1) return this; // Fast-path
        if (clearCofactor) return clearCofactor(Point, this);
        return this.multiplyUnsafe(CURVE.h);
      }
    }, {
      key: "toRawBytes",
      value: function toRawBytes() {
        var isCompressed = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : true;
        this.assertValidity();
        return toBytes(Point, this, isCompressed);
      }
    }, {
      key: "toHex",
      value: function toHex() {
        var isCompressed = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : true;
        return bytesToHex(this.toRawBytes(isCompressed));
      }
    }], [{
      key: "fromAffine",
      value: function fromAffine(p) {
        var _ref = p || {},
          x = _ref.x,
          y = _ref.y;
        if (!p || !Fp.isValid(x) || !Fp.isValid(y)) throw new Error('invalid affine point');
        if (p instanceof Point) throw new Error('projective point not allowed');
        var is0 = function is0(i) {
          return Fp.eql(i, Fp.ZERO);
        };
        // fromAffine(x:0, y:0) would produce (x:0, y:0, z:1), but we need (x:0, y:1, z:0)
        if (is0(x) && is0(y)) return Point.ZERO;
        return new Point(x, y, Fp.ONE);
      }
    }, {
      key: "normalizeZ",
      value: function normalizeZ(points) {
        var toInv = Fp.invertBatch(points.map(function (p) {
          return p.pz;
        }));
        return points.map(function (p, i) {
          return p.toAffine(toInv[i]);
        }).map(Point.fromAffine);
      }
      /**
       * Converts hash string or Uint8Array to Point.
       * @param hex short/long ECDSA hex
       */
    }, {
      key: "fromHex",
      value: function fromHex(hex) {
        var P = Point.fromAffine(fromBytes(ensureBytes('pointHex', hex)));
        P.assertValidity();
        return P;
      }
      // Multiplies generator point by privateKey.
    }, {
      key: "fromPrivateKey",
      value: function fromPrivateKey(privateKey) {
        return Point.BASE.multiply(normPrivateKeyToScalar(privateKey));
      }
    }]);
  }();
  Point.BASE = new Point(CURVE.Gx, CURVE.Gy, Fp.ONE);
  Point.ZERO = new Point(Fp.ZERO, Fp.ONE, Fp.ZERO);
  var _bits = CURVE.nBitLength;
  var wnaf = wNAF(Point, CURVE.endo ? Math.ceil(_bits / 2) : _bits);
  // Validate if generator point is on curve
  return {
    CURVE: CURVE,
    ProjectivePoint: Point,
    normPrivateKeyToScalar: normPrivateKeyToScalar,
    weierstrassEquation: weierstrassEquation,
    isWithinCurveOrder: isWithinCurveOrder
  };
}
function validateOpts(curve) {
  var opts = validateBasic(curve);
  validateObject(opts, {
    hash: 'hash',
    hmac: 'function',
    randomBytes: 'function'
  }, {
    bits2int: 'function',
    bits2int_modN: 'function',
    lowS: 'boolean'
  });
  return Object.freeze(_objectSpread2({
    lowS: true
  }, opts));
}
function weierstrass(curveDef) {
  var CURVE = validateOpts(curveDef);
  var Fp = CURVE.Fp,
    CURVE_ORDER = CURVE.n;
  var compressedLen = Fp.BYTES + 1; // e.g. 33 for 32
  var uncompressedLen = 2 * Fp.BYTES + 1; // e.g. 65 for 32
  function isValidFieldElement(num) {
    return _0n < num && num < Fp.ORDER; // 0 is banned since it's not invertible FE
  }
  function modN(a) {
    return mod(a, CURVE_ORDER);
  }
  function invN(a) {
    return invert(a, CURVE_ORDER);
  }
  var _weierstrassPoints = weierstrassPoints(_objectSpread2(_objectSpread2({}, CURVE), {}, {
      toBytes: function toBytes(_c, point, isCompressed) {
        var a = point.toAffine();
        var x = Fp.toBytes(a.x);
        var cat = concatBytes;
        if (isCompressed) {
          return cat(Uint8Array.from([point.hasEvenY() ? 0x02 : 0x03]), x);
        } else {
          return cat(Uint8Array.from([0x04]), x, Fp.toBytes(a.y));
        }
      },
      fromBytes: function fromBytes(bytes) {
        var len = bytes.length;
        var head = bytes[0];
        var tail = bytes.subarray(1);
        // this.assertValidity() is done inside of fromHex
        if (len === compressedLen && (head === 0x02 || head === 0x03)) {
          var x = bytesToNumberBE(tail);
          if (!isValidFieldElement(x)) throw new Error('Point is not on curve');
          var y2 = weierstrassEquation(x); // y² = x³ + ax + b
          var y = Fp.sqrt(y2); // y = y² ^ (p+1)/4
          var isYOdd = (y & _1n$1) === _1n$1;
          // ECDSA
          var isHeadOdd = (head & 1) === 1;
          if (isHeadOdd !== isYOdd) y = Fp.neg(y);
          return {
            x: x,
            y: y
          };
        } else if (len === uncompressedLen && head === 0x04) {
          var _x = Fp.fromBytes(tail.subarray(0, Fp.BYTES));
          var _y = Fp.fromBytes(tail.subarray(Fp.BYTES, 2 * Fp.BYTES));
          return {
            x: _x,
            y: _y
          };
        } else {
          throw new Error("Point of length ".concat(len, " was invalid. Expected ").concat(compressedLen, " compressed bytes or ").concat(uncompressedLen, " uncompressed bytes"));
        }
      }
    })),
    Point = _weierstrassPoints.ProjectivePoint,
    normPrivateKeyToScalar = _weierstrassPoints.normPrivateKeyToScalar,
    weierstrassEquation = _weierstrassPoints.weierstrassEquation,
    isWithinCurveOrder = _weierstrassPoints.isWithinCurveOrder;
  var numToNByteStr = function numToNByteStr(num) {
    return bytesToHex(numberToBytesBE(num, CURVE.nByteLength));
  };
  function isBiggerThanHalfOrder(number) {
    var HALF = CURVE_ORDER >> _1n$1;
    return number > HALF;
  }
  function normalizeS(s) {
    return isBiggerThanHalfOrder(s) ? modN(-s) : s;
  }
  // slice bytes num
  var slcNum = function slcNum(b, from, to) {
    return bytesToNumberBE(b.slice(from, to));
  };
  /**
   * ECDSA signature with its (r, s) properties. Supports DER & compact representations.
   */
  var Signature = /*#__PURE__*/function () {
    function Signature(r, s, recovery) {
      _classCallCheck(this, Signature);
      this.r = r;
      this.s = s;
      this.recovery = recovery;
      this.assertValidity();
    }
    // pair (bytes of r, bytes of s)
    return _createClass(Signature, [{
      key: "assertValidity",
      value: function assertValidity() {
        // can use assertGE here
        if (!isWithinCurveOrder(this.r)) throw new Error('r must be 0 < r < CURVE.n');
        if (!isWithinCurveOrder(this.s)) throw new Error('s must be 0 < s < CURVE.n');
      }
    }, {
      key: "addRecoveryBit",
      value: function addRecoveryBit(recovery) {
        return new Signature(this.r, this.s, recovery);
      }
    }, {
      key: "recoverPublicKey",
      value: function recoverPublicKey(msgHash) {
        var r = this.r,
          s = this.s,
          rec = this.recovery;
        var h = bits2int_modN(ensureBytes('msgHash', msgHash)); // Truncate hash
        if (rec == null || ![0, 1, 2, 3].includes(rec)) throw new Error('recovery id invalid');
        var radj = rec === 2 || rec === 3 ? r + CURVE.n : r;
        if (radj >= Fp.ORDER) throw new Error('recovery id 2 or 3 invalid');
        var prefix = (rec & 1) === 0 ? '02' : '03';
        var R = Point.fromHex(prefix + numToNByteStr(radj));
        var ir = invN(radj); // r^-1
        var u1 = modN(-h * ir); // -hr^-1
        var u2 = modN(s * ir); // sr^-1
        var Q = Point.BASE.multiplyAndAddUnsafe(R, u1, u2); // (sr^-1)R-(hr^-1)G = -(hr^-1)G + (sr^-1)
        if (!Q) throw new Error('point at infinify'); // unsafe is fine: no priv data leaked
        Q.assertValidity();
        return Q;
      }
      // Signatures should be low-s, to prevent malleability.
    }, {
      key: "hasHighS",
      value: function hasHighS() {
        return isBiggerThanHalfOrder(this.s);
      }
    }, {
      key: "normalizeS",
      value: function normalizeS() {
        return this.hasHighS() ? new Signature(this.r, modN(-this.s), this.recovery) : this;
      }
      // DER-encoded
    }, {
      key: "toDERRawBytes",
      value: function toDERRawBytes() {
        return hexToBytes(this.toDERHex());
      }
    }, {
      key: "toDERHex",
      value: function toDERHex() {
        return DER.hexFromSig({
          r: this.r,
          s: this.s
        });
      }
      // padded bytes of r, then padded bytes of s
    }, {
      key: "toCompactRawBytes",
      value: function toCompactRawBytes() {
        return hexToBytes(this.toCompactHex());
      }
    }, {
      key: "toCompactHex",
      value: function toCompactHex() {
        return numToNByteStr(this.r) + numToNByteStr(this.s);
      }
    }], [{
      key: "fromCompact",
      value: function fromCompact(hex) {
        var l = CURVE.nByteLength;
        hex = ensureBytes('compactSignature', hex, l * 2);
        return new Signature(slcNum(hex, 0, l), slcNum(hex, l, 2 * l));
      }
      // DER encoded ECDSA signature
      // https://bitcoin.stackexchange.com/questions/57644/what-are-the-parts-of-a-bitcoin-transaction-input-script
    }, {
      key: "fromDER",
      value: function fromDER(hex) {
        var _DER$toSig = DER.toSig(ensureBytes('DER', hex)),
          r = _DER$toSig.r,
          s = _DER$toSig.s;
        return new Signature(r, s);
      }
    }]);
  }();
  var utils = {
    isValidPrivateKey: function isValidPrivateKey(privateKey) {
      try {
        normPrivateKeyToScalar(privateKey);
        return true;
      } catch (error) {
        return false;
      }
    },
    normPrivateKeyToScalar: normPrivateKeyToScalar,
    /**
     * Produces cryptographically secure private key from random of size
     * (groupLen + ceil(groupLen / 2)) with modulo bias being negligible.
     */
    randomPrivateKey: function randomPrivateKey() {
      var length = getMinHashLength(CURVE.n);
      return mapHashToField(CURVE.randomBytes(length), CURVE.n);
    },
    /**
     * Creates precompute table for an arbitrary EC point. Makes point "cached".
     * Allows to massively speed-up `point.multiply(scalar)`.
     * @returns cached point
     * @example
     * const fast = utils.precompute(8, ProjectivePoint.fromHex(someonesPubKey));
     * fast.multiply(privKey); // much faster ECDH now
     */
    precompute: function precompute() {
      var windowSize = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : 8;
      var point = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : Point.BASE;
      point._setWindowSize(windowSize);
      point.multiply(BigInt(3)); // 3 is arbitrary, just need any number here
      return point;
    }
  };
  /**
   * Computes public key for a private key. Checks for validity of the private key.
   * @param privateKey private key
   * @param isCompressed whether to return compact (default), or full key
   * @returns Public key, full when isCompressed=false; short when isCompressed=true
   */
  function getPublicKey(privateKey) {
    var isCompressed = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : true;
    return Point.fromPrivateKey(privateKey).toRawBytes(isCompressed);
  }
  /**
   * Quick and dirty check for item being public key. Does not validate hex, or being on-curve.
   */
  function isProbPub(item) {
    var arr = item instanceof Uint8Array;
    var str = typeof item === 'string';
    var len = (arr || str) && item.length;
    if (arr) return len === compressedLen || len === uncompressedLen;
    if (str) return len === 2 * compressedLen || len === 2 * uncompressedLen;
    if (item instanceof Point) return true;
    return false;
  }
  /**
   * ECDH (Elliptic Curve Diffie Hellman).
   * Computes shared public key from private key and public key.
   * Checks: 1) private key validity 2) shared key is on-curve.
   * Does NOT hash the result.
   * @param privateA private key
   * @param publicB different public key
   * @param isCompressed whether to return compact (default), or full key
   * @returns shared public key
   */
  function getSharedSecret(privateA, publicB) {
    var isCompressed = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : true;
    if (isProbPub(privateA)) throw new Error('first arg must be private key');
    if (!isProbPub(publicB)) throw new Error('second arg must be public key');
    var b = Point.fromHex(publicB); // check for being on-curve
    return b.multiply(normPrivateKeyToScalar(privateA)).toRawBytes(isCompressed);
  }
  // RFC6979: ensure ECDSA msg is X bytes and < N. RFC suggests optional truncating via bits2octets.
  // FIPS 186-4 4.6 suggests the leftmost min(nBitLen, outLen) bits, which matches bits2int.
  // bits2int can produce res>N, we can do mod(res, N) since the bitLen is the same.
  // int2octets can't be used; pads small msgs with 0: unacceptatble for trunc as per RFC vectors
  var bits2int = CURVE.bits2int || function (bytes) {
    // For curves with nBitLength % 8 !== 0: bits2octets(bits2octets(m)) !== bits2octets(m)
    // for some cases, since bytes.length * 8 is not actual bitLength.
    var num = bytesToNumberBE(bytes); // check for == u8 done here
    var delta = bytes.length * 8 - CURVE.nBitLength; // truncate to nBitLength leftmost bits
    return delta > 0 ? num >> BigInt(delta) : num;
  };
  var bits2int_modN = CURVE.bits2int_modN || function (bytes) {
    return modN(bits2int(bytes)); // can't use bytesToNumberBE here
  };
  // NOTE: pads output with zero as per spec
  var ORDER_MASK = bitMask(CURVE.nBitLength);
  /**
   * Converts to bytes. Checks if num in `[0..ORDER_MASK-1]` e.g.: `[0..2^256-1]`.
   */
  function int2octets(num) {
    if (typeof num !== 'bigint') throw new Error('bigint expected');
    if (!(_0n <= num && num < ORDER_MASK)) throw new Error("bigint expected < 2^".concat(CURVE.nBitLength));
    // works with order, can have different size than numToField!
    return numberToBytesBE(num, CURVE.nByteLength);
  }
  // Steps A, D of RFC6979 3.2
  // Creates RFC6979 seed; converts msg/privKey to numbers.
  // Used only in sign, not in verify.
  // NOTE: we cannot assume here that msgHash has same amount of bytes as curve order, this will be wrong at least for P521.
  // Also it can be bigger for P224 + SHA256
  function prepSig(msgHash, privateKey) {
    var opts = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : defaultSigOpts;
    if (['recovered', 'canonical'].some(function (k) {
      return k in opts;
    })) throw new Error('sign() legacy options not supported');
    var hash = CURVE.hash,
      randomBytes = CURVE.randomBytes;
    var lowS = opts.lowS,
      prehash = opts.prehash,
      ent = opts.extraEntropy; // generates low-s sigs by default
    if (lowS == null) lowS = true; // RFC6979 3.2: we skip step A, because we already provide hash
    msgHash = ensureBytes('msgHash', msgHash);
    if (prehash) msgHash = ensureBytes('prehashed msgHash', hash(msgHash));
    // We can't later call bits2octets, since nested bits2int is broken for curves
    // with nBitLength % 8 !== 0. Because of that, we unwrap it here as int2octets call.
    // const bits2octets = (bits) => int2octets(bits2int_modN(bits))
    var h1int = bits2int_modN(msgHash);
    var d = normPrivateKeyToScalar(privateKey); // validate private key, convert to bigint
    var seedArgs = [int2octets(d), int2octets(h1int)];
    // extraEntropy. RFC6979 3.6: additional k' (optional).
    if (ent != null) {
      // K = HMAC_K(V || 0x00 || int2octets(x) || bits2octets(h1) || k')
      var e = ent === true ? randomBytes(Fp.BYTES) : ent; // generate random bytes OR pass as-is
      seedArgs.push(ensureBytes('extraEntropy', e)); // check for being bytes
    }
    var seed = concatBytes.apply(ut, seedArgs); // Step D of RFC6979 3.2
    var m = h1int; // NOTE: no need to call bits2int second time here, it is inside truncateHash!
    // Converts signature params into point w r/s, checks result for validity.
    function k2sig(kBytes) {
      // RFC 6979 Section 3.2, step 3: k = bits2int(T)
      var k = bits2int(kBytes); // Cannot use fields methods, since it is group element
      if (!isWithinCurveOrder(k)) return; // Important: all mod() calls here must be done over N
      var ik = invN(k); // k^-1 mod n
      var q = Point.BASE.multiply(k).toAffine(); // q = Gk
      var r = modN(q.x); // r = q.x mod n
      if (r === _0n) return;
      // Can use scalar blinding b^-1(bm + bdr) where b ∈ [1,q−1] according to
      // https://tches.iacr.org/index.php/TCHES/article/view/7337/6509. We've decided against it:
      // a) dependency on CSPRNG b) 15% slowdown c) doesn't really help since bigints are not CT
      var s = modN(ik * modN(m + r * d)); // Not using blinding here
      if (s === _0n) return;
      var recovery = (q.x === r ? 0 : 2) | Number(q.y & _1n$1); // recovery bit (2 or 3, when q.x > n)
      var normS = s;
      if (lowS && isBiggerThanHalfOrder(s)) {
        normS = normalizeS(s); // if lowS was passed, ensure s is always
        recovery ^= 1; // // in the bottom half of N
      }
      return new Signature(r, normS, recovery); // use normS, not s
    }
    return {
      seed: seed,
      k2sig: k2sig
    };
  }
  var defaultSigOpts = {
    lowS: CURVE.lowS,
    prehash: false
  };
  var defaultVerOpts = {
    lowS: CURVE.lowS,
    prehash: false
  };
  /**
   * Signs message hash with a private key.
   * ```
   * sign(m, d, k) where
   *   (x, y) = G × k
   *   r = x mod n
   *   s = (m + dr)/k mod n
   * ```
   * @param msgHash NOT message. msg needs to be hashed to `msgHash`, or use `prehash`.
   * @param privKey private key
   * @param opts lowS for non-malleable sigs. extraEntropy for mixing randomness into k. prehash will hash first arg.
   * @returns signature with recovery param
   */
  function sign(msgHash, privKey) {
    var opts = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : defaultSigOpts;
    var _prepSig = prepSig(msgHash, privKey, opts),
      seed = _prepSig.seed,
      k2sig = _prepSig.k2sig; // Steps A, D of RFC6979 3.2.
    var C = CURVE;
    var drbg = createHmacDrbg(C.hash.outputLen, C.nByteLength, C.hmac);
    return drbg(seed, k2sig); // Steps B, C, D, E, F, G
  }
  // Enable precomputes. Slows down first publicKey computation by 20ms.
  Point.BASE._setWindowSize(8);
  // utils.precompute(8, ProjectivePoint.BASE)
  /**
   * Verifies a signature against message hash and public key.
   * Rejects lowS signatures by default: to override,
   * specify option `{lowS: false}`. Implements section 4.1.4 from https://www.secg.org/sec1-v2.pdf:
   *
   * ```
   * verify(r, s, h, P) where
   *   U1 = hs^-1 mod n
   *   U2 = rs^-1 mod n
   *   R = U1⋅G - U2⋅P
   *   mod(R.x, n) == r
   * ```
   */
  function verify(signature, msgHash, publicKey) {
    var _Point$BASE$multiplyA;
    var opts = arguments.length > 3 && arguments[3] !== undefined ? arguments[3] : defaultVerOpts;
    var sg = signature;
    msgHash = ensureBytes('msgHash', msgHash);
    publicKey = ensureBytes('publicKey', publicKey);
    if ('strict' in opts) throw new Error('options.strict was renamed to lowS');
    var lowS = opts.lowS,
      prehash = opts.prehash;
    var _sig = undefined;
    var P;
    try {
      if (typeof sg === 'string' || sg instanceof Uint8Array) {
        // Signature can be represented in 2 ways: compact (2*nByteLength) & DER (variable-length).
        // Since DER can also be 2*nByteLength bytes, we check for it first.
        try {
          _sig = Signature.fromDER(sg);
        } catch (derError) {
          if (!(derError instanceof DER.Err)) throw derError;
          _sig = Signature.fromCompact(sg);
        }
      } else if (_typeof(sg) === 'object' && typeof sg.r === 'bigint' && typeof sg.s === 'bigint') {
        var _r = sg.r,
          _s = sg.s;
        _sig = new Signature(_r, _s);
      } else {
        throw new Error('PARSE');
      }
      P = Point.fromHex(publicKey);
    } catch (error) {
      if (error.message === 'PARSE') throw new Error("signature must be Signature instance, Uint8Array or hex string");
      return false;
    }
    if (lowS && _sig.hasHighS()) return false;
    if (prehash) msgHash = CURVE.hash(msgHash);
    var _sig2 = _sig,
      r = _sig2.r,
      s = _sig2.s;
    var h = bits2int_modN(msgHash); // Cannot use fields methods, since it is group element
    var is = invN(s); // s^-1
    var u1 = modN(h * is); // u1 = hs^-1 mod n
    var u2 = modN(r * is); // u2 = rs^-1 mod n
    var R = (_Point$BASE$multiplyA = Point.BASE.multiplyAndAddUnsafe(P, u1, u2)) === null || _Point$BASE$multiplyA === void 0 ? void 0 : _Point$BASE$multiplyA.toAffine(); // R = u1⋅G + u2⋅P
    if (!R) return false;
    var v = modN(R.x);
    return v === r;
  }
  return {
    CURVE: CURVE,
    getPublicKey: getPublicKey,
    getSharedSecret: getSharedSecret,
    sign: sign,
    verify: verify,
    ProjectivePoint: Point,
    Signature: Signature,
    utils: utils
  };
}

// connects noble-curves to noble-hashes
function getHash(hash) {
  return {
    hash: hash,
    hmac: function hmac$1(key) {
      for (var _len = arguments.length, msgs = new Array(_len > 1 ? _len - 1 : 0), _key = 1; _key < _len; _key++) {
        msgs[_key - 1] = arguments[_key];
      }
      return hmac(hash, key, concatBytes$1.apply(void 0, msgs));
    },
    randomBytes: randomBytes
  };
}
function createCurve(curveDef, defHash) {
  var create = function create(hash) {
    return weierstrass(_objectSpread2(_objectSpread2({}, curveDef), getHash(hash)));
  };
  return Object.freeze(_objectSpread2(_objectSpread2({}, create(defHash)), {}, {
    create: create
  }));
}

/*! noble-curves - MIT License (c) 2022 Paul Miller (paulmillr.com) */
var secp256k1P = BigInt('0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f');
var secp256k1N = BigInt('0xfffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141');
var _1n = BigInt(1);
var _2n = BigInt(2);
var divNearest = function divNearest(a, b) {
  return (a + b / _2n) / b;
};
/**
 * √n = n^((p+1)/4) for fields p = 3 mod 4. We unwrap the loop and multiply bit-by-bit.
 * (P+1n/4n).toString(2) would produce bits [223x 1, 0, 22x 1, 4x 0, 11, 00]
 */
function sqrtMod(y) {
  var P = secp256k1P;
  // prettier-ignore
  var _3n = BigInt(3),
    _6n = BigInt(6),
    _11n = BigInt(11),
    _22n = BigInt(22);
  // prettier-ignore
  var _23n = BigInt(23),
    _44n = BigInt(44),
    _88n = BigInt(88);
  var b2 = y * y * y % P; // x^3, 11
  var b3 = b2 * b2 * y % P; // x^7
  var b6 = pow2(b3, _3n, P) * b3 % P;
  var b9 = pow2(b6, _3n, P) * b3 % P;
  var b11 = pow2(b9, _2n, P) * b2 % P;
  var b22 = pow2(b11, _11n, P) * b11 % P;
  var b44 = pow2(b22, _22n, P) * b22 % P;
  var b88 = pow2(b44, _44n, P) * b44 % P;
  var b176 = pow2(b88, _88n, P) * b88 % P;
  var b220 = pow2(b176, _44n, P) * b44 % P;
  var b223 = pow2(b220, _3n, P) * b3 % P;
  var t1 = pow2(b223, _23n, P) * b22 % P;
  var t2 = pow2(t1, _6n, P) * b2 % P;
  var root = pow2(t2, _2n, P);
  if (!Fp.eql(Fp.sqr(root), y)) throw new Error('Cannot find square root');
  return root;
}
var Fp = Field(secp256k1P, undefined, undefined, {
  sqrt: sqrtMod
});
var secp256k1 = createCurve({
  a: BigInt(0),
  b: BigInt(7),
  Fp: Fp,
  n: secp256k1N,
  // Base point (x, y) aka generator point
  Gx: BigInt('55066263022277343669578718895168534326250603453777594175500187360389116729240'),
  Gy: BigInt('32670510020758816978083085130507043184471273380659243275938904335757337482424'),
  h: BigInt(1),
  lowS: true,
  /**
   * secp256k1 belongs to Koblitz curves: it has efficiently computable endomorphism.
   * Endomorphism uses 2x less RAM, speeds up precomputation by 2x and ECDH / key recovery by 20%.
   * For precomputed wNAF it trades off 1/2 init time & 1/3 ram for 20% perf hit.
   * Explanation: https://gist.github.com/paulmillr/eb670806793e84df628a7c434a873066
   */
  endo: {
    beta: BigInt('0x7ae96a2b657c07106e64479eac3434e99cf0497512f58995c1396c28719501ee'),
    splitScalar: function splitScalar(k) {
      var n = secp256k1N;
      var a1 = BigInt('0x3086d221a7d46bcde86c90e49284eb15');
      var b1 = -_1n * BigInt('0xe4437ed6010e88286f547fa90abfe4c3');
      var a2 = BigInt('0x114ca50f7a8e2f3f657c1108d9d44cfd8');
      var b2 = a1;
      var POW_2_128 = BigInt('0x100000000000000000000000000000000'); // (2n**128n).toString(16)
      var c1 = divNearest(b2 * k, n);
      var c2 = divNearest(-b1 * k, n);
      var k1 = mod(k - c1 * a1 - c2 * a2, n);
      var k2 = mod(-c1 * b1 - c2 * b2, n);
      var k1neg = k1 > POW_2_128;
      var k2neg = k2 > POW_2_128;
      if (k1neg) k1 = n - k1;
      if (k2neg) k2 = n - k2;
      if (k1 > POW_2_128 || k2 > POW_2_128) {
        throw new Error('splitScalar: Endomorphism failed, k=' + k);
      }
      return {
        k1neg: k1neg,
        k1: k1,
        k2neg: k2neg,
        k2: k2
      };
    }
  }
}, sha256);
// Schnorr signatures are superior to ECDSA from above. Below is Schnorr-specific BIP0340 code.
// https://github.com/bitcoin/bips/blob/master/bip-0340.mediawiki
BigInt(0);
secp256k1.ProjectivePoint;

/**
 *  A constant for the zero hash.
 *
 *  (**i.e.** ``"0x0000000000000000000000000000000000000000000000000000000000000000"``)
 */
var ZeroHash = "0x0000000000000000000000000000000000000000000000000000000000000000";

// NFKC (composed)             // (decomposed)
/**
 *  A constant for the ether symbol (normalized using NFKC).
 *
 *  (**i.e.** ``"\\u039e"``)
 */
/**
 *  A constant for the [[link-eip-191]] personal message prefix.
 *
 *  (**i.e.** ``"\\x19Ethereum Signed Message:\\n"``)
 */
var MessagePrefix = "\x19Ethereum Signed Message:\n";

// Constants
var BN_0$1 = BigInt(0);
var BN_1 = BigInt(1);
var BN_2 = BigInt(2);
var BN_27 = BigInt(27);
var BN_28 = BigInt(28);
var BN_35 = BigInt(35);
var _guard = {};
function toUint256(value) {
  return zeroPadValue(toBeArray(value), 32);
}
/**
 *  A Signature  @TODO
 *
 *
 *  @_docloc: api/crypto:Signing
 */
var _r2 = /*#__PURE__*/new WeakMap();
var _s = /*#__PURE__*/new WeakMap();
var _v2 = /*#__PURE__*/new WeakMap();
var _networkV = /*#__PURE__*/new WeakMap();
var Signature = /*#__PURE__*/function () {
  /**
   *  @private
   */
  function Signature(guard, r, s, v) {
    _classCallCheck(this, Signature);
    _classPrivateFieldInitSpec(this, _r2, void 0);
    _classPrivateFieldInitSpec(this, _s, void 0);
    _classPrivateFieldInitSpec(this, _v2, void 0);
    _classPrivateFieldInitSpec(this, _networkV, void 0);
    assertPrivate(guard, _guard, "Signature");
    _classPrivateFieldSet2(_r2, this, r);
    _classPrivateFieldSet2(_s, this, s);
    _classPrivateFieldSet2(_v2, this, v);
    _classPrivateFieldSet2(_networkV, this, null);
  }
  return _createClass(Signature, [{
    key: "r",
    get:
    /**
     *  The ``r`` value for a signautre.
     *
     *  This represents the ``x`` coordinate of a "reference" or
     *  challenge point, from which the ``y`` can be computed.
     */
    function get() {
      return _classPrivateFieldGet2(_r2, this);
    },
    set: function set(value) {
      assertArgument(dataLength(value) === 32, "invalid r", "value", value);
      _classPrivateFieldSet2(_r2, this, hexlify(value));
    }
    /**
     *  The ``s`` value for a signature.
     */
  }, {
    key: "s",
    get: function get() {
      return _classPrivateFieldGet2(_s, this);
    },
    set: function set(_value) {
      assertArgument(dataLength(_value) === 32, "invalid s", "value", _value);
      var value = hexlify(_value);
      assertArgument(parseInt(value.substring(0, 3)) < 8, "non-canonical s", "value", value);
      _classPrivateFieldSet2(_s, this, value);
    }
    /**
     *  The ``v`` value for a signature.
     *
     *  Since a given ``x`` value for ``r`` has two possible values for
     *  its correspondin ``y``, the ``v`` indicates which of the two ``y``
     *  values to use.
     *
     *  It is normalized to the values ``27`` or ``28`` for legacy
     *  purposes.
     */
  }, {
    key: "v",
    get: function get() {
      return _classPrivateFieldGet2(_v2, this);
    },
    set: function set(value) {
      var v = getNumber(value, "value");
      assertArgument(v === 27 || v === 28, "invalid v", "v", value);
      _classPrivateFieldSet2(_v2, this, v);
    }
    /**
     *  The EIP-155 ``v`` for legacy transactions. For non-legacy
     *  transactions, this value is ``null``.
     */
  }, {
    key: "networkV",
    get: function get() {
      return _classPrivateFieldGet2(_networkV, this);
    }
    /**
     *  The chain ID for EIP-155 legacy transactions. For non-legacy
     *  transactions, this value is ``null``.
     */
  }, {
    key: "legacyChainId",
    get: function get() {
      var v = this.networkV;
      if (v == null) {
        return null;
      }
      return Signature.getChainId(v);
    }
    /**
     *  The ``yParity`` for the signature.
     *
     *  See ``v`` for more details on how this value is used.
     */
  }, {
    key: "yParity",
    get: function get() {
      return this.v === 27 ? 0 : 1;
    }
    /**
     *  The [[link-eip-2098]] compact representation of the ``yParity``
     *  and ``s`` compacted into a single ``bytes32``.
     */
  }, {
    key: "yParityAndS",
    get: function get() {
      // The EIP-2098 compact representation
      var yParityAndS = getBytes(this.s);
      if (this.yParity) {
        yParityAndS[0] |= 0x80;
      }
      return hexlify(yParityAndS);
    }
    /**
     *  The [[link-eip-2098]] compact representation.
     */
  }, {
    key: "compactSerialized",
    get: function get() {
      return concat([this.r, this.yParityAndS]);
    }
    /**
     *  The serialized representation.
     */
  }, {
    key: "serialized",
    get: function get() {
      return concat([this.r, this.s, this.yParity ? "0x1c" : "0x1b"]);
    }
  }, {
    key: Symbol["for"]('nodejs.util.inspect.custom'),
    value: function value() {
      return "Signature { r: \"".concat(this.r, "\", s: \"").concat(this.s, "\", yParity: ").concat(this.yParity, ", networkV: ").concat(this.networkV, " }");
    }
    /**
     *  Returns a new identical [[Signature]].
     */
  }, {
    key: "clone",
    value: function clone() {
      var clone = new Signature(_guard, this.r, this.s, this.v);
      if (this.networkV) {
        _classPrivateFieldSet2(_networkV, clone, this.networkV);
      }
      return clone;
    }
    /**
     *  Returns a representation that is compatible with ``JSON.stringify``.
     */
  }, {
    key: "toJSON",
    value: function toJSON() {
      var networkV = this.networkV;
      return {
        _type: "signature",
        networkV: networkV != null ? networkV.toString() : null,
        r: this.r,
        s: this.s,
        v: this.v
      };
    }
    /**
     *  Compute the chain ID from the ``v`` in a legacy EIP-155 transactions.
     *
     *  @example:
     *    Signature.getChainId(45)
     *    //_result:
     *
     *    Signature.getChainId(46)
     *    //_result:
     */
  }], [{
    key: "getChainId",
    value: function getChainId(v) {
      var bv = getBigInt(v, "v");
      // The v is not an EIP-155 v, so it is the unspecified chain ID
      if (bv == BN_27 || bv == BN_28) {
        return BN_0$1;
      }
      // Bad value for an EIP-155 v
      assertArgument(bv >= BN_35, "invalid EIP-155 v", "v", v);
      return (bv - BN_35) / BN_2;
    }
    /**
     *  Compute the ``v`` for a chain ID for a legacy EIP-155 transactions.
     *
     *  Legacy transactions which use [[link-eip-155]] hijack the ``v``
     *  property to include the chain ID.
     *
     *  @example:
     *    Signature.getChainIdV(5, 27)
     *    //_result:
     *
     *    Signature.getChainIdV(5, 28)
     *    //_result:
     *
     */
  }, {
    key: "getChainIdV",
    value: function getChainIdV(chainId, v) {
      return getBigInt(chainId) * BN_2 + BigInt(35 + v - 27);
    }
    /**
     *  Compute the normalized legacy transaction ``v`` from a ``yParirty``,
     *  a legacy transaction ``v`` or a legacy [[link-eip-155]] transaction.
     *
     *  @example:
     *    // The values 0 and 1 imply v is actually yParity
     *    Signature.getNormalizedV(0)
     *    //_result:
     *
     *    // Legacy non-EIP-1559 transaction (i.e. 27 or 28)
     *    Signature.getNormalizedV(27)
     *    //_result:
     *
     *    // Legacy EIP-155 transaction (i.e. >= 35)
     *    Signature.getNormalizedV(46)
     *    //_result:
     *
     *    // Invalid values throw
     *    Signature.getNormalizedV(5)
     *    //_error:
     */
  }, {
    key: "getNormalizedV",
    value: function getNormalizedV(v) {
      var bv = getBigInt(v);
      if (bv === BN_0$1 || bv === BN_27) {
        return 27;
      }
      if (bv === BN_1 || bv === BN_28) {
        return 28;
      }
      assertArgument(bv >= BN_35, "invalid v", "v", v);
      // Otherwise, EIP-155 v means odd is 27 and even is 28
      return bv & BN_1 ? 27 : 28;
    }
    /**
     *  Creates a new [[Signature]].
     *
     *  If no %%sig%% is provided, a new [[Signature]] is created
     *  with default values.
     *
     *  If %%sig%% is a string, it is parsed.
     */
  }, {
    key: "from",
    value: function from(sig) {
      function assertError(check, message) {
        assertArgument(check, message, "signature", sig);
      }
      if (sig == null) {
        return new Signature(_guard, ZeroHash, ZeroHash, 27);
      }
      if (typeof sig === "string") {
        var bytes = getBytes(sig, "signature");
        if (bytes.length === 64) {
          var _r3 = hexlify(bytes.slice(0, 32));
          var _s2 = bytes.slice(32, 64);
          var _v3 = _s2[0] & 0x80 ? 28 : 27;
          _s2[0] &= 0x7f;
          return new Signature(_guard, _r3, hexlify(_s2), _v3);
        }
        if (bytes.length === 65) {
          var _r4 = hexlify(bytes.slice(0, 32));
          var _s3 = bytes.slice(32, 64);
          assertError((_s3[0] & 0x80) === 0, "non-canonical s");
          var _v4 = Signature.getNormalizedV(bytes[64]);
          return new Signature(_guard, _r4, hexlify(_s3), _v4);
        }
        assertError(false, "invalid raw signature length");
      }
      if (sig instanceof Signature) {
        return sig.clone();
      }
      // Get r
      var _r = sig.r;
      assertError(_r != null, "missing r");
      var r = toUint256(_r);
      // Get s; by any means necessary (we check consistency below)
      var s = function (s, yParityAndS) {
        if (s != null) {
          return toUint256(s);
        }
        if (yParityAndS != null) {
          assertError(isHexString(yParityAndS, 32), "invalid yParityAndS");
          var _bytes = getBytes(yParityAndS);
          _bytes[0] &= 0x7f;
          return hexlify(_bytes);
        }
        assertError(false, "missing s");
      }(sig.s, sig.yParityAndS);
      assertError((getBytes(s)[0] & 0x80) == 0, "non-canonical s");
      // Get v; by any means necessary (we check consistency below)
      var _ref = function (_v, yParityAndS, yParity) {
          if (_v != null) {
            var _v5 = getBigInt(_v);
            return {
              networkV: _v5 >= BN_35 ? _v5 : undefined,
              v: Signature.getNormalizedV(_v5)
            };
          }
          if (yParityAndS != null) {
            assertError(isHexString(yParityAndS, 32), "invalid yParityAndS");
            return {
              v: getBytes(yParityAndS)[0] & 0x80 ? 28 : 27
            };
          }
          if (yParity != null) {
            switch (getNumber(yParity, "sig.yParity")) {
              case 0:
                return {
                  v: 27
                };
              case 1:
                return {
                  v: 28
                };
            }
            assertError(false, "invalid yParity");
          }
          assertError(false, "missing v");
        }(sig.v, sig.yParityAndS, sig.yParity),
        networkV = _ref.networkV,
        v = _ref.v;
      var result = new Signature(_guard, r, s, v);
      if (networkV) {
        _classPrivateFieldSet2(_networkV, result, networkV);
      }
      // If multiple of v, yParity, yParityAndS we given, check they match
      assertError(sig.yParity == null || getNumber(sig.yParity, "sig.yParity") === result.yParity, "yParity mismatch");
      assertError(sig.yParityAndS == null || sig.yParityAndS === result.yParityAndS, "yParityAndS mismatch");
      return result;
    }
  }]);
}();

/**
 *  A **SigningKey** provides high-level access to the elliptic curve
 *  cryptography (ECC) operations and key management.
 */
var _privateKey = /*#__PURE__*/new WeakMap();
var SigningKey = /*#__PURE__*/function () {
  /**
   *  Creates a new **SigningKey** for %%privateKey%%.
   */
  function SigningKey(privateKey) {
    _classCallCheck(this, SigningKey);
    _classPrivateFieldInitSpec(this, _privateKey, void 0);
    assertArgument(dataLength(privateKey) === 32, "invalid private key", "privateKey", "[REDACTED]");
    _classPrivateFieldSet2(_privateKey, this, hexlify(privateKey));
  }
  /**
   *  The private key.
   */
  return _createClass(SigningKey, [{
    key: "privateKey",
    get: function get() {
      return _classPrivateFieldGet2(_privateKey, this);
    }
    /**
     *  The uncompressed public key.
     *
     * This will always begin with the prefix ``0x04`` and be 132
     * characters long (the ``0x`` prefix and 130 hexadecimal nibbles).
     */
  }, {
    key: "publicKey",
    get: function get() {
      return SigningKey.computePublicKey(_classPrivateFieldGet2(_privateKey, this));
    }
    /**
     *  The compressed public key.
     *
     *  This will always begin with either the prefix ``0x02`` or ``0x03``
     *  and be 68 characters long (the ``0x`` prefix and 33 hexadecimal
     *  nibbles)
     */
  }, {
    key: "compressedPublicKey",
    get: function get() {
      return SigningKey.computePublicKey(_classPrivateFieldGet2(_privateKey, this), true);
    }
    /**
     *  Return the signature of the signed %%digest%%.
     */
  }, {
    key: "sign",
    value: function sign(digest) {
      assertArgument(dataLength(digest) === 32, "invalid digest length", "digest", digest);
      var sig = secp256k1.sign(getBytesCopy(digest), getBytesCopy(_classPrivateFieldGet2(_privateKey, this)), {
        lowS: true
      });
      return Signature.from({
        r: toBeHex(sig.r, 32),
        s: toBeHex(sig.s, 32),
        v: sig.recovery ? 0x1c : 0x1b
      });
    }
    /**
     *  Returns the [[link-wiki-ecdh]] shared secret between this
     *  private key and the %%other%% key.
     *
     *  The %%other%% key may be any type of key, a raw public key,
     *  a compressed/uncompressed pubic key or aprivate key.
     *
     *  Best practice is usually to use a cryptographic hash on the
     *  returned value before using it as a symetric secret.
     *
     *  @example:
     *    sign1 = new SigningKey(id("some-secret-1"))
     *    sign2 = new SigningKey(id("some-secret-2"))
     *
     *    // Notice that privA.computeSharedSecret(pubB)...
     *    sign1.computeSharedSecret(sign2.publicKey)
     *    //_result:
     *
     *    // ...is equal to privB.computeSharedSecret(pubA).
     *    sign2.computeSharedSecret(sign1.publicKey)
     *    //_result:
     */
  }, {
    key: "computeSharedSecret",
    value: function computeSharedSecret(other) {
      var pubKey = SigningKey.computePublicKey(other);
      return hexlify(secp256k1.getSharedSecret(getBytesCopy(_classPrivateFieldGet2(_privateKey, this)), getBytes(pubKey), false));
    }
    /**
     *  Compute the public key for %%key%%, optionally %%compressed%%.
     *
     *  The %%key%% may be any type of key, a raw public key, a
     *  compressed/uncompressed public key or private key.
     *
     *  @example:
     *    sign = new SigningKey(id("some-secret"));
     *
     *    // Compute the uncompressed public key for a private key
     *    SigningKey.computePublicKey(sign.privateKey)
     *    //_result:
     *
     *    // Compute the compressed public key for a private key
     *    SigningKey.computePublicKey(sign.privateKey, true)
     *    //_result:
     *
     *    // Compute the uncompressed public key
     *    SigningKey.computePublicKey(sign.publicKey, false);
     *    //_result:
     *
     *    // Compute the Compressed a public key
     *    SigningKey.computePublicKey(sign.publicKey, true);
     *    //_result:
     */
  }], [{
    key: "computePublicKey",
    value: function computePublicKey(key, compressed) {
      var bytes = getBytes(key, "key");
      // private key
      if (bytes.length === 32) {
        var pubKey = secp256k1.getPublicKey(bytes, !!compressed);
        return hexlify(pubKey);
      }
      // raw public key; use uncompressed key with 0x04 prefix
      if (bytes.length === 64) {
        var pub = new Uint8Array(65);
        pub[0] = 0x04;
        pub.set(bytes, 1);
        bytes = pub;
      }
      var point = secp256k1.ProjectivePoint.fromHex(bytes);
      return hexlify(point.toRawBytes(compressed));
    }
    /**
     *  Returns the public key for the private key which produced the
     *  %%signature%% for the given %%digest%%.
     *
     *  @example:
     *    key = new SigningKey(id("some-secret"))
     *    digest = id("hello world")
     *    sig = key.sign(digest)
     *
     *    // Notice the signer public key...
     *    key.publicKey
     *    //_result:
     *
     *    // ...is equal to the recovered public key
     *    SigningKey.recoverPublicKey(digest, sig)
     *    //_result:
     *
     */
  }, {
    key: "recoverPublicKey",
    value: function recoverPublicKey(digest, signature) {
      assertArgument(dataLength(digest) === 32, "invalid digest length", "digest", digest);
      var sig = Signature.from(signature);
      var secpSig = secp256k1.Signature.fromCompact(getBytesCopy(concat([sig.r, sig.s])));
      secpSig = secpSig.addRecoveryBit(sig.yParity);
      var pubKey = secpSig.recoverPublicKey(getBytesCopy(digest));
      assertArgument(pubKey != null, "invalid signautre for digest", "signature", signature);
      return "0x" + pubKey.toHex(false);
    }
    /**
     *  Returns the point resulting from adding the ellipic curve points
     *  %%p0%% and %%p1%%.
     *
     *  This is not a common function most developers should require, but
     *  can be useful for certain privacy-specific techniques.
     *
     *  For example, it is used by [[HDNodeWallet]] to compute child
     *  addresses from parent public keys and chain codes.
     */
  }, {
    key: "addPoints",
    value: function addPoints(p0, p1, compressed) {
      var pub0 = secp256k1.ProjectivePoint.fromHex(SigningKey.computePublicKey(p0).substring(2));
      var pub1 = secp256k1.ProjectivePoint.fromHex(SigningKey.computePublicKey(p1).substring(2));
      return "0x" + pub0.add(pub1).toHex(!!compressed);
    }
  }]);
}();

var BN_0 = BigInt(0);
var BN_36 = BigInt(36);
function getChecksumAddress(address) {
  //    if (!isHexString(address, 20)) {
  //        logger.throwArgumentError("invalid address", "address", address);
  //    }
  address = address.toLowerCase();
  var chars = address.substring(2).split("");
  var expanded = new Uint8Array(40);
  for (var i = 0; i < 40; i++) {
    expanded[i] = chars[i].charCodeAt(0);
  }
  var hashed = getBytes(keccak256(expanded));
  for (var _i = 0; _i < 40; _i += 2) {
    if (hashed[_i >> 1] >> 4 >= 8) {
      chars[_i] = chars[_i].toUpperCase();
    }
    if ((hashed[_i >> 1] & 0x0f) >= 8) {
      chars[_i + 1] = chars[_i + 1].toUpperCase();
    }
  }
  return "0x" + chars.join("");
}
// See: https://en.wikipedia.org/wiki/International_Bank_Account_Number
// Create lookup table
var ibanLookup = {};
for (var i = 0; i < 10; i++) {
  ibanLookup[String(i)] = String(i);
}
for (var _i2 = 0; _i2 < 26; _i2++) {
  ibanLookup[String.fromCharCode(65 + _i2)] = String(10 + _i2);
}
// How many decimal digits can we process? (for 64-bit float, this is 15)
// i.e. Math.floor(Math.log10(Number.MAX_SAFE_INTEGER));
var safeDigits = 15;
function ibanChecksum(address) {
  address = address.toUpperCase();
  address = address.substring(4) + address.substring(0, 2) + "00";
  var expanded = address.split("").map(function (c) {
    return ibanLookup[c];
  }).join("");
  // Javascript can handle integers safely up to 15 (decimal) digits
  while (expanded.length >= safeDigits) {
    var block = expanded.substring(0, safeDigits);
    expanded = parseInt(block, 10) % 97 + expanded.substring(block.length);
  }
  var checksum = String(98 - parseInt(expanded, 10) % 97);
  while (checksum.length < 2) {
    checksum = "0" + checksum;
  }
  return checksum;
}
var Base36 = function () {
  var result = {};
  for (var _i3 = 0; _i3 < 36; _i3++) {
    var key = "0123456789abcdefghijklmnopqrstuvwxyz"[_i3];
    result[key] = BigInt(_i3);
  }
  return result;
}();
function fromBase36(value) {
  value = value.toLowerCase();
  var result = BN_0;
  for (var _i4 = 0; _i4 < value.length; _i4++) {
    result = result * BN_36 + Base36[value[_i4]];
  }
  return result;
}
/**
 *  Returns a normalized and checksumed address for %%address%%.
 *  This accepts non-checksum addresses, checksum addresses and
 *  [[getIcapAddress]] formats.
 *
 *  The checksum in Ethereum uses the capitalization (upper-case
 *  vs lower-case) of the characters within an address to encode
 *  its checksum, which offers, on average, a checksum of 15-bits.
 *
 *  If %%address%% contains both upper-case and lower-case, it is
 *  assumed to already be a checksum address and its checksum is
 *  validated, and if the address fails its expected checksum an
 *  error is thrown.
 *
 *  If you wish the checksum of %%address%% to be ignore, it should
 *  be converted to lower-case (i.e. ``.toLowercase()``) before
 *  being passed in. This should be a very rare situation though,
 *  that you wish to bypass the safegaurds in place to protect
 *  against an address that has been incorrectly copied from another
 *  source.
 *
 *  @example:
 *    // Adds the checksum (via upper-casing specific letters)
 *    getAddress("0x8ba1f109551bd432803012645ac136ddd64dba72")
 *    //_result:
 *
 *    // Converts ICAP address and adds checksum
 *    getAddress("XE65GB6LDNXYOFTX0NSV3FUWKOWIXAMJK36");
 *    //_result:
 *
 *    // Throws an error if an address contains mixed case,
 *    // but the checksum fails
 *    getAddress("0x8Ba1f109551bD432803012645Ac136ddd64DBA72")
 *    //_error:
 */
function getAddress(address) {
  assertArgument(typeof address === "string", "invalid address", "address", address);
  if (address.match(/^(0x)?[0-9a-fA-F]{40}$/)) {
    // Missing the 0x prefix
    if (!address.startsWith("0x")) {
      address = "0x" + address;
    }
    var result = getChecksumAddress(address);
    // It is a checksummed address with a bad checksum
    assertArgument(!address.match(/([A-F].*[a-f])|([a-f].*[A-F])/) || result === address, "bad address checksum", "address", address);
    return result;
  }
  // Maybe ICAP? (we only support direct mode)
  if (address.match(/^XE[0-9]{2}[0-9A-Za-z]{30,31}$/)) {
    // It is an ICAP address with a bad checksum
    assertArgument(address.substring(2, 4) === ibanChecksum(address), "bad icap checksum", "address", address);
    var _result = fromBase36(address.substring(4)).toString(16);
    while (_result.length < 40) {
      _result = "0" + _result;
    }
    return getChecksumAddress("0x" + _result);
  }
  assertArgument(false, "invalid address", "address", address);
}

/**
 *  Returns the address for the %%key%%.
 *
 *  The key may be any standard form of public key or a private key.
 */
function computeAddress(key) {
  var pubkey;
  if (typeof key === "string") {
    pubkey = SigningKey.computePublicKey(key, false);
  } else {
    pubkey = key.publicKey;
  }
  return getAddress(keccak256("0x" + pubkey.substring(4)).substring(26));
}
/**
 *  Returns the recovered address for the private key that was
 *  used to sign %%digest%% that resulted in %%signature%%.
 */
function recoverAddress(digest, signature) {
  return computeAddress(SigningKey.recoverPublicKey(digest, signature));
}

/**
 *  Computes the [[link-eip-191]] personal-sign message digest to sign.
 *
 *  This prefixes the message with [[MessagePrefix]] and the decimal length
 *  of %%message%% and computes the [[keccak256]] digest.
 *
 *  If %%message%% is a string, it is converted to its UTF-8 bytes
 *  first. To compute the digest of a [[DataHexString]], it must be converted
 *  to [bytes](getBytes).
 *
 *  @example:
 *    hashMessage("Hello World")
 *    //_result:
 *
 *    // Hashes the SIX (6) string characters, i.e.
 *    // [ "0", "x", "4", "2", "4", "3" ]
 *    hashMessage("0x4243")
 *    //_result:
 *
 *    // Hashes the TWO (2) bytes [ 0x42, 0x43 ]...
 *    hashMessage(getBytes("0x4243"))
 *    //_result:
 *
 *    // ...which is equal to using data
 *    hashMessage(new Uint8Array([ 0x42, 0x43 ]))
 *    //_result:
 *
 */
function hashMessage(message) {
  if (typeof message === "string") {
    message = toUtf8Bytes(message);
  }
  return keccak256(concat([toUtf8Bytes(MessagePrefix), toUtf8Bytes(String(message.length)), message]));
}
/**
 *  Return the address of the private key that produced
 *  the signature %%sig%% during signing for %%message%%.
 */
function verifyMessage(message, sig) {
  var digest = hashMessage(message);
  return recoverAddress(digest, sig);
}

var BeforeAvalancheAuthenticate = function BeforeAvalancheAuthenticate(ctx, logger, nk, data) {
  var _a = JSON.parse(data.account.id),
    chain = _a.chain,
    message = _a.message,
    address = _a.address,
    signature = _a.signature;
  if (chain !== Chain.Avalanche) {
    return data;
  }
  var verified = verifyMessage(message, signature) === address;
  if (!verified) {
    logger.error("Signature verification failed");
    return null;
  } else {
    return _assign(_assign({}, data), {
      account: {
        id: address
      }
    });
  }
};

function rpcHealthcheck(ctx, logger, nk, payload) {
  return JSON.stringify({
    status: "OK"
  });
}

var InitModule = function InitModule(ctx, logger, nk, initializer) {
  initializer.registerBeforeAuthenticateCustom(BeforeAvalancheAuthenticate);
  initializer.registerRpc("typescript_healthcheck", rpcHealthcheck);
  logger.info("Hello World!");
};
!InitModule && InitModule.bind(null);
