import require$$0 from 'fs';
import require$$1 from 'url';
import require$$2 from 'child_process';
import require$$3 from 'http';
import require$$4 from 'https';
import require$$0$1 from 'stream';
import require$$0$2 from 'zlib';
import require$$0$3 from 'buffer';
import require$$5 from 'crypto';
import require$$0$4 from 'events';
import require$$3$1 from 'net';
import require$$4$1 from 'tls';

function _mergeNamespaces(n, m) {
    m.forEach(function (e) {
        e && typeof e !== 'string' && !Array.isArray(e) && Object.keys(e).forEach(function (k) {
            if (k !== 'default' && !(k in n)) {
                var d = Object.getOwnPropertyDescriptor(e, k);
                Object.defineProperty(n, k, d.get ? d : {
                    enumerable: true,
                    get: function () { return e[k]; }
                });
            }
        });
    });
    return Object.freeze(n);
}

function rpcHealthcheck(ctx, logger, nk, payload) {
  return JSON.stringify({
    status: "OK"
  });
}

function _arrayLikeToArray(r, a) {
  (null == a || a > r.length) && (a = r.length);
  for (var e = 0, n = Array(a); e < a; e++) n[e] = r[e];
  return n;
}
function _arrayWithHoles(r) {
  if (Array.isArray(r)) return r;
}
function _assertThisInitialized(e) {
  if (void 0 === e) throw new ReferenceError("this hasn't been initialised - super() hasn't been called");
  return e;
}
function _callSuper(t, o, e) {
  return o = _getPrototypeOf(o), _possibleConstructorReturn(t, _isNativeReflectConstruct() ? Reflect.construct(o, e || [], _getPrototypeOf(t).constructor) : o.apply(t, e));
}
function _classCallCheck(a, n) {
  if (!(a instanceof n)) throw new TypeError("Cannot call a class as a function");
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
function _createForOfIteratorHelper(r, e) {
  var t = "undefined" != typeof Symbol && r[Symbol.iterator] || r["@@iterator"];
  if (!t) {
    if (Array.isArray(r) || (t = _unsupportedIterableToArray(r)) || e) {
      t && (r = t);
      var n = 0,
        F = function () {};
      return {
        s: F,
        n: function () {
          return n >= r.length ? {
            done: !0
          } : {
            done: !1,
            value: r[n++]
          };
        },
        e: function (r) {
          throw r;
        },
        f: F
      };
    }
    throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.");
  }
  var o,
    a = !0,
    u = !1;
  return {
    s: function () {
      t = t.call(r);
    },
    n: function () {
      var r = t.next();
      return a = r.done, r;
    },
    e: function (r) {
      u = !0, o = r;
    },
    f: function () {
      try {
        a || null == t.return || t.return();
      } finally {
        if (u) throw o;
      }
    }
  };
}
function _defineProperty(e, r, t) {
  return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, {
    value: t,
    enumerable: !0,
    configurable: !0,
    writable: !0
  }) : e[r] = t, e;
}
function _get() {
  return _get = "undefined" != typeof Reflect && Reflect.get ? Reflect.get.bind() : function (e, t, r) {
    var p = _superPropBase(e, t);
    if (p) {
      var n = Object.getOwnPropertyDescriptor(p, t);
      return n.get ? n.get.call(arguments.length < 3 ? e : r) : n.value;
    }
  }, _get.apply(null, arguments);
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
function _superPropBase(t, o) {
  for (; !{}.hasOwnProperty.call(t, o) && null !== (t = _getPrototypeOf(t)););
  return t;
}
function _superPropGet(t, e, o, r) {
  var p = _get(_getPrototypeOf(1 & r ? t.prototype : t), e, o);
  return 2 & r && "function" == typeof p ? function (t) {
    return p.apply(o, t);
  } : p;
}
function _toPrimitive(t, r) {
  if ("object" != typeof t || !t) return t;
  var e = t[Symbol.toPrimitive];
  if (void 0 !== e) {
    var i = e.call(t, r || "default");
    if ("object" != typeof i) return i;
    throw new TypeError("@@toPrimitive must return a primitive value.");
  }
  return ("string" === r ? String : Number)(t);
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

var PACKET_TYPES = Object.create(null); // no Map = no polyfill
PACKET_TYPES["open"] = "0";
PACKET_TYPES["close"] = "1";
PACKET_TYPES["ping"] = "2";
PACKET_TYPES["pong"] = "3";
PACKET_TYPES["message"] = "4";
PACKET_TYPES["upgrade"] = "5";
PACKET_TYPES["noop"] = "6";
var PACKET_TYPES_REVERSE = Object.create(null);
Object.keys(PACKET_TYPES).forEach(function (key) {
  PACKET_TYPES_REVERSE[PACKET_TYPES[key]] = key;
});
var ERROR_PACKET = {
  type: "error",
  data: "parser error"
};

var encodePacket = function encodePacket(_ref, supportsBinary, callback) {
  var type = _ref.type,
    data = _ref.data;
  if (data instanceof ArrayBuffer || ArrayBuffer.isView(data)) {
    return callback(supportsBinary ? data : "b" + toBuffer$3(data, true).toString("base64"));
  }
  // plain string
  return callback(PACKET_TYPES[type] + (data || ""));
};
var toBuffer$3 = function toBuffer(data, forceBufferConversion) {
  if (Buffer.isBuffer(data) || data instanceof Uint8Array && !forceBufferConversion) {
    return data;
  } else if (data instanceof ArrayBuffer) {
    return Buffer.from(data);
  } else {
    return Buffer.from(data.buffer, data.byteOffset, data.byteLength);
  }
};
var TEXT_ENCODER;
function encodePacketToBinary(packet, callback) {
  if (packet.data instanceof ArrayBuffer || ArrayBuffer.isView(packet.data)) {
    return callback(toBuffer$3(packet.data, false));
  }
  encodePacket(packet, true, function (encoded) {
    if (!TEXT_ENCODER) {
      // lazily created for compatibility with Node.js 10
      TEXT_ENCODER = new TextEncoder();
    }
    callback(TEXT_ENCODER.encode(encoded));
  });
}

var decodePacket = function decodePacket(encodedPacket, binaryType) {
  if (typeof encodedPacket !== "string") {
    return {
      type: "message",
      data: mapBinary(encodedPacket, binaryType)
    };
  }
  var type = encodedPacket.charAt(0);
  if (type === "b") {
    var buffer = Buffer.from(encodedPacket.substring(1), "base64");
    return {
      type: "message",
      data: mapBinary(buffer, binaryType)
    };
  }
  if (!PACKET_TYPES_REVERSE[type]) {
    return ERROR_PACKET;
  }
  return encodedPacket.length > 1 ? {
    type: PACKET_TYPES_REVERSE[type],
    data: encodedPacket.substring(1)
  } : {
    type: PACKET_TYPES_REVERSE[type]
  };
};
var mapBinary = function mapBinary(data, binaryType) {
  switch (binaryType) {
    case "arraybuffer":
      if (data instanceof ArrayBuffer) {
        // from WebSocket & binaryType "arraybuffer"
        return data;
      } else if (Buffer.isBuffer(data)) {
        // from HTTP long-polling
        return data.buffer.slice(data.byteOffset, data.byteOffset + data.byteLength);
      } else {
        // from WebTransport (Uint8Array)
        return data.buffer;
      }
    case "nodebuffer":
    default:
      if (Buffer.isBuffer(data)) {
        // from HTTP long-polling or WebSocket & binaryType "nodebuffer" (default)
        return data;
      } else {
        // from WebTransport (Uint8Array)
        return Buffer.from(data);
      }
  }
};

var SEPARATOR = String.fromCharCode(30); // see https://en.wikipedia.org/wiki/Delimiter#ASCII_delimited_text
var encodePayload = function encodePayload(packets, callback) {
  // some packets may be added to the array while encoding, so the initial length must be saved
  var length = packets.length;
  var encodedPackets = new Array(length);
  var count = 0;
  packets.forEach(function (packet, i) {
    // force base64 encoding for binary packets
    encodePacket(packet, false, function (encodedPacket) {
      encodedPackets[i] = encodedPacket;
      if (++count === length) {
        callback(encodedPackets.join(SEPARATOR));
      }
    });
  });
};
var decodePayload = function decodePayload(encodedPayload, binaryType) {
  var encodedPackets = encodedPayload.split(SEPARATOR);
  var packets = [];
  for (var i = 0; i < encodedPackets.length; i++) {
    var decodedPacket = decodePacket(encodedPackets[i], binaryType);
    packets.push(decodedPacket);
    if (decodedPacket.type === "error") {
      break;
    }
  }
  return packets;
};
function createPacketEncoderStream() {
  return new TransformStream({
    transform: function transform(packet, controller) {
      encodePacketToBinary(packet, function (encodedPacket) {
        var payloadLength = encodedPacket.length;
        var header;
        // inspired by the WebSocket format: https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_servers#decoding_payload_length
        if (payloadLength < 126) {
          header = new Uint8Array(1);
          new DataView(header.buffer).setUint8(0, payloadLength);
        } else if (payloadLength < 65536) {
          header = new Uint8Array(3);
          var view = new DataView(header.buffer);
          view.setUint8(0, 126);
          view.setUint16(1, payloadLength);
        } else {
          header = new Uint8Array(9);
          var _view = new DataView(header.buffer);
          _view.setUint8(0, 127);
          _view.setBigUint64(1, BigInt(payloadLength));
        }
        // first bit indicates whether the payload is plain text (0) or binary (1)
        if (packet.data && typeof packet.data !== "string") {
          header[0] |= 0x80;
        }
        controller.enqueue(header);
        controller.enqueue(encodedPacket);
      });
    }
  });
}
var TEXT_DECODER;
function totalLength(chunks) {
  return chunks.reduce(function (acc, chunk) {
    return acc + chunk.length;
  }, 0);
}
function concatChunks(chunks, size) {
  if (chunks[0].length === size) {
    return chunks.shift();
  }
  var buffer = new Uint8Array(size);
  var j = 0;
  for (var i = 0; i < size; i++) {
    buffer[i] = chunks[0][j++];
    if (j === chunks[0].length) {
      chunks.shift();
      j = 0;
    }
  }
  if (chunks.length && j < chunks[0].length) {
    chunks[0] = chunks[0].slice(j);
  }
  return buffer;
}
function createPacketDecoderStream(maxPayload, binaryType) {
  if (!TEXT_DECODER) {
    TEXT_DECODER = new TextDecoder();
  }
  var chunks = [];
  var state = 0 /* State.READ_HEADER */;
  var expectedLength = -1;
  var isBinary = false;
  return new TransformStream({
    transform: function transform(chunk, controller) {
      chunks.push(chunk);
      while (true) {
        if (state === 0 /* State.READ_HEADER */) {
          if (totalLength(chunks) < 1) {
            break;
          }
          var header = concatChunks(chunks, 1);
          isBinary = (header[0] & 0x80) === 0x80;
          expectedLength = header[0] & 0x7f;
          if (expectedLength < 126) {
            state = 3 /* State.READ_PAYLOAD */;
          } else if (expectedLength === 126) {
            state = 1 /* State.READ_EXTENDED_LENGTH_16 */;
          } else {
            state = 2 /* State.READ_EXTENDED_LENGTH_64 */;
          }
        } else if (state === 1 /* State.READ_EXTENDED_LENGTH_16 */) {
          if (totalLength(chunks) < 2) {
            break;
          }
          var headerArray = concatChunks(chunks, 2);
          expectedLength = new DataView(headerArray.buffer, headerArray.byteOffset, headerArray.length).getUint16(0);
          state = 3 /* State.READ_PAYLOAD */;
        } else if (state === 2 /* State.READ_EXTENDED_LENGTH_64 */) {
          if (totalLength(chunks) < 8) {
            break;
          }
          var _headerArray = concatChunks(chunks, 8);
          var view = new DataView(_headerArray.buffer, _headerArray.byteOffset, _headerArray.length);
          var n = view.getUint32(0);
          if (n > Math.pow(2, 53 - 32) - 1) {
            // the maximum safe integer in JavaScript is 2^53 - 1
            controller.enqueue(ERROR_PACKET);
            break;
          }
          expectedLength = n * Math.pow(2, 32) + view.getUint32(4);
          state = 3 /* State.READ_PAYLOAD */;
        } else {
          if (totalLength(chunks) < expectedLength) {
            break;
          }
          var data = concatChunks(chunks, expectedLength);
          controller.enqueue(decodePacket(isBinary ? data : TEXT_DECODER.decode(data), binaryType));
          state = 0 /* State.READ_HEADER */;
        }
        if (expectedLength === 0 || expectedLength > maxPayload) {
          controller.enqueue(ERROR_PACKET);
          break;
        }
      }
    }
  });
}
var protocol$1 = 4;

/**
 * Initialize a new `Emitter`.
 *
 * @api public
 */

function Emitter(obj) {
  if (obj) return mixin(obj);
}

/**
 * Mixin the emitter properties.
 *
 * @param {Object} obj
 * @return {Object}
 * @api private
 */

function mixin(obj) {
  for (var key in Emitter.prototype) {
    obj[key] = Emitter.prototype[key];
  }
  return obj;
}

/**
 * Listen on the given `event` with `fn`.
 *
 * @param {String} event
 * @param {Function} fn
 * @return {Emitter}
 * @api public
 */

Emitter.prototype.on = Emitter.prototype.addEventListener = function (event, fn) {
  this._callbacks = this._callbacks || {};
  (this._callbacks['$' + event] = this._callbacks['$' + event] || []).push(fn);
  return this;
};

/**
 * Adds an `event` listener that will be invoked a single
 * time then automatically removed.
 *
 * @param {String} event
 * @param {Function} fn
 * @return {Emitter}
 * @api public
 */

Emitter.prototype.once = function (event, fn) {
  function on() {
    this.off(event, on);
    fn.apply(this, arguments);
  }
  on.fn = fn;
  this.on(event, on);
  return this;
};

/**
 * Remove the given callback for `event` or all
 * registered callbacks.
 *
 * @param {String} event
 * @param {Function} fn
 * @return {Emitter}
 * @api public
 */

Emitter.prototype.off = Emitter.prototype.removeListener = Emitter.prototype.removeAllListeners = Emitter.prototype.removeEventListener = function (event, fn) {
  this._callbacks = this._callbacks || {};

  // all
  if (0 == arguments.length) {
    this._callbacks = {};
    return this;
  }

  // specific event
  var callbacks = this._callbacks['$' + event];
  if (!callbacks) return this;

  // remove all handlers
  if (1 == arguments.length) {
    delete this._callbacks['$' + event];
    return this;
  }

  // remove specific handler
  var cb;
  for (var i = 0; i < callbacks.length; i++) {
    cb = callbacks[i];
    if (cb === fn || cb.fn === fn) {
      callbacks.splice(i, 1);
      break;
    }
  }

  // Remove event specific arrays for event types that no
  // one is subscribed for to avoid memory leak.
  if (callbacks.length === 0) {
    delete this._callbacks['$' + event];
  }
  return this;
};

/**
 * Emit `event` with the given args.
 *
 * @param {String} event
 * @param {Mixed} ...
 * @return {Emitter}
 */

Emitter.prototype.emit = function (event) {
  this._callbacks = this._callbacks || {};
  var args = new Array(arguments.length - 1),
    callbacks = this._callbacks['$' + event];
  for (var i = 1; i < arguments.length; i++) {
    args[i - 1] = arguments[i];
  }
  if (callbacks) {
    callbacks = callbacks.slice(0);
    for (var i = 0, len = callbacks.length; i < len; ++i) {
      callbacks[i].apply(this, args);
    }
  }
  return this;
};

// alias used for reserved events (protected method)
Emitter.prototype.emitReserved = Emitter.prototype.emit;

/**
 * Return array of callbacks for `event`.
 *
 * @param {String} event
 * @return {Array}
 * @api public
 */

Emitter.prototype.listeners = function (event) {
  this._callbacks = this._callbacks || {};
  return this._callbacks['$' + event] || [];
};

/**
 * Check if this emitter has `event` handlers.
 *
 * @param {String} event
 * @return {Boolean}
 * @api public
 */

Emitter.prototype.hasListeners = function (event) {
  return !!this.listeners(event).length;
};

var globalThisShim = global;

function pick(obj) {
  for (var _len = arguments.length, attr = new Array(_len > 1 ? _len - 1 : 0), _key = 1; _key < _len; _key++) {
    attr[_key - 1] = arguments[_key];
  }
  return attr.reduce(function (acc, k) {
    if (obj.hasOwnProperty(k)) {
      acc[k] = obj[k];
    }
    return acc;
  }, {});
}
// Keep a reference to the real timeout functions so they can be used when overridden
var NATIVE_SET_TIMEOUT = globalThisShim.setTimeout;
var NATIVE_CLEAR_TIMEOUT = globalThisShim.clearTimeout;
function installTimerFunctions(obj, opts) {
  if (opts.useNativeTimers) {
    obj.setTimeoutFn = NATIVE_SET_TIMEOUT.bind(globalThisShim);
    obj.clearTimeoutFn = NATIVE_CLEAR_TIMEOUT.bind(globalThisShim);
  } else {
    obj.setTimeoutFn = globalThisShim.setTimeout.bind(globalThisShim);
    obj.clearTimeoutFn = globalThisShim.clearTimeout.bind(globalThisShim);
  }
}
// base64 encoded buffers are about 33% bigger (https://en.wikipedia.org/wiki/Base64)
var BASE64_OVERHEAD = 1.33;
// we could also have used `new Blob([obj]).size`, but it isn't supported in IE9
function byteLength(obj) {
  if (typeof obj === "string") {
    return utf8Length(obj);
  }
  // arraybuffer or blob
  return Math.ceil((obj.byteLength || obj.size) * BASE64_OVERHEAD);
}
function utf8Length(str) {
  var c = 0,
    length = 0;
  for (var i = 0, l = str.length; i < l; i++) {
    c = str.charCodeAt(i);
    if (c < 0x80) {
      length += 1;
    } else if (c < 0x800) {
      length += 2;
    } else if (c < 0xd800 || c >= 0xe000) {
      length += 3;
    } else {
      i++;
      length += 4;
    }
  }
  return length;
}

// imported from https://github.com/galkn/querystring
/**
 * Compiles a querystring
 * Returns string representation of the object
 *
 * @param {Object}
 * @api private
 */
function encode$1(obj) {
  var str = '';
  for (var i in obj) {
    if (obj.hasOwnProperty(i)) {
      if (str.length) str += '&';
      str += encodeURIComponent(i) + '=' + encodeURIComponent(obj[i]);
    }
  }
  return str;
}
/**
 * Parses a simple querystring into an object
 *
 * @param {String} qs
 * @api private
 */
function decode(qs) {
  var qry = {};
  var pairs = qs.split('&');
  for (var i = 0, l = pairs.length; i < l; i++) {
    var pair = pairs[i].split('=');
    qry[decodeURIComponent(pair[0])] = decodeURIComponent(pair[1]);
  }
  return qry;
}

var TransportError = /*#__PURE__*/function (_Error) {
  function TransportError(reason, description, context) {
    var _this;
    _classCallCheck(this, TransportError);
    _this = _callSuper(this, TransportError, [reason]);
    _this.description = description;
    _this.context = context;
    _this.type = "TransportError";
    return _this;
  }
  _inherits(TransportError, _Error);
  return _createClass(TransportError);
}(/*#__PURE__*/_wrapNativeSuper(Error));
var Transport = /*#__PURE__*/function (_Emitter) {
  /**
   * Transport abstract constructor.
   *
   * @param {Object} opts - options
   * @protected
   */
  function Transport(opts) {
    var _this2;
    _classCallCheck(this, Transport);
    _this2 = _callSuper(this, Transport);
    _this2.writable = false;
    installTimerFunctions(_this2, opts);
    _this2.opts = opts;
    _this2.query = opts.query;
    _this2.socket = opts.socket;
    return _this2;
  }
  /**
   * Emits an error.
   *
   * @param {String} reason
   * @param description
   * @param context - the error context
   * @return {Transport} for chaining
   * @protected
   */
  _inherits(Transport, _Emitter);
  return _createClass(Transport, [{
    key: "onError",
    value: function onError(reason, description, context) {
      _superPropGet(Transport, "emitReserved", this, 3)(["error", new TransportError(reason, description, context)]);
      return this;
    }
    /**
     * Opens the transport.
     */
  }, {
    key: "open",
    value: function open() {
      this.readyState = "opening";
      this.doOpen();
      return this;
    }
    /**
     * Closes the transport.
     */
  }, {
    key: "close",
    value: function close() {
      if (this.readyState === "opening" || this.readyState === "open") {
        this.doClose();
        this.onClose();
      }
      return this;
    }
    /**
     * Sends multiple packets.
     *
     * @param {Array} packets
     */
  }, {
    key: "send",
    value: function send(packets) {
      if (this.readyState === "open") {
        this.write(packets);
      }
    }
    /**
     * Called upon open
     *
     * @protected
     */
  }, {
    key: "onOpen",
    value: function onOpen() {
      this.readyState = "open";
      this.writable = true;
      _superPropGet(Transport, "emitReserved", this, 3)(["open"]);
    }
    /**
     * Called with data.
     *
     * @param {String} data
     * @protected
     */
  }, {
    key: "onData",
    value: function onData(data) {
      var packet = decodePacket(data, this.socket.binaryType);
      this.onPacket(packet);
    }
    /**
     * Called with a decoded packet.
     *
     * @protected
     */
  }, {
    key: "onPacket",
    value: function onPacket(packet) {
      _superPropGet(Transport, "emitReserved", this, 3)(["packet", packet]);
    }
    /**
     * Called upon close.
     *
     * @protected
     */
  }, {
    key: "onClose",
    value: function onClose(details) {
      this.readyState = "closed";
      _superPropGet(Transport, "emitReserved", this, 3)(["close", details]);
    }
    /**
     * Pauses the transport, in order not to lose packets during an upgrade.
     *
     * @param onPause
     */
  }, {
    key: "pause",
    value: function pause(onPause) {}
  }, {
    key: "createUri",
    value: function createUri(schema) {
      var query = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : {};
      return schema + "://" + this._hostname() + this._port() + this.opts.path + this._query(query);
    }
  }, {
    key: "_hostname",
    value: function _hostname() {
      var hostname = this.opts.hostname;
      return hostname.indexOf(":") === -1 ? hostname : "[" + hostname + "]";
    }
  }, {
    key: "_port",
    value: function _port() {
      if (this.opts.port && (this.opts.secure && Number(this.opts.port !== 443) || !this.opts.secure && Number(this.opts.port) !== 80)) {
        return ":" + this.opts.port;
      } else {
        return "";
      }
    }
  }, {
    key: "_query",
    value: function _query(query) {
      var encodedQuery = encode$1(query);
      return encodedQuery.length ? "?" + encodedQuery : "";
    }
  }]);
}(Emitter);

// imported from https://github.com/unshiftio/yeast

var alphabet = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_'.split(''),
  length = 64,
  map = {};
var seed = 0,
  i = 0,
  prev;
/**
 * Return a string representing the specified number.
 *
 * @param {Number} num The number to convert.
 * @returns {String} The string representation of the number.
 * @api public
 */
function encode(num) {
  var encoded = '';
  do {
    encoded = alphabet[num % length] + encoded;
    num = Math.floor(num / length);
  } while (num > 0);
  return encoded;
}
/**
 * Yeast: A tiny growing id generator.
 *
 * @returns {String} A unique id.
 * @api public
 */
function yeast() {
  var now = encode(+new Date());
  if (now !== prev) return seed = 0, prev = now;
  return now + '.' + encode(seed++);
}
//
// Map each character to its index.
//
for (; i < length; i++) map[alphabet[i]] = i;

function getDefaultExportFromCjs (x) {
	return x && x.__esModule && Object.prototype.hasOwnProperty.call(x, 'default') ? x['default'] : x;
}

/**
 * Wrapper for built-in http.js to emulate the browser XMLHttpRequest object.
 *
 * This can be used with JS designed for browsers to improve reuse of code and
 * allow the use of existing libraries.
 *
 * Usage: include("XMLHttpRequest.js") and use XMLHttpRequest per W3C specs.
 *
 * @author Dan DeFelippi <dan@driverdan.com>
 * @contributor David Ellis <d.f.ellis@ieee.org>
 * @license MIT
 */
var fs = require$$0;
var Url = require$$1;
var spawn = require$$2.spawn;

/**
 * Module exports.
 */

var XMLHttpRequest_1 = XMLHttpRequest;

// backwards-compat
XMLHttpRequest.XMLHttpRequest = XMLHttpRequest;

/**
 * `XMLHttpRequest` constructor.
 *
 * Supported options for the `opts` object are:
 *
 *  - `agent`: An http.Agent instance; http.globalAgent may be used; if 'undefined', agent usage is disabled
 *
 * @param {Object} opts optional "options" object
 */

function XMLHttpRequest(opts) {

  opts = opts || {};

  /**
   * Private variables
   */
  var self = this;
  var http = require$$3;
  var https = require$$4;

  // Holds http.js objects
  var request;
  var response;

  // Request settings
  var settings = {};

  // Disable header blacklist.
  // Not part of XHR specs.
  var disableHeaderCheck = false;

  // Set some default headers
  var defaultHeaders = {
    "User-Agent": "node-XMLHttpRequest",
    "Accept": "*/*"
  };
  var headers = Object.assign({}, defaultHeaders);

  // These headers are not user setable.
  // The following are allowed but banned in the spec:
  // * user-agent
  var forbiddenRequestHeaders = ["accept-charset", "accept-encoding", "access-control-request-headers", "access-control-request-method", "connection", "content-length", "content-transfer-encoding", "cookie", "cookie2", "date", "expect", "host", "keep-alive", "origin", "referer", "te", "trailer", "transfer-encoding", "upgrade", "via"];

  // These request methods are not allowed
  var forbiddenRequestMethods = ["TRACE", "TRACK", "CONNECT"];

  // Send flag
  var sendFlag = false;
  // Error flag, used when errors occur or abort is called
  var errorFlag = false;
  var abortedFlag = false;

  // Event listeners
  var listeners = {};

  /**
   * Constants
   */

  this.UNSENT = 0;
  this.OPENED = 1;
  this.HEADERS_RECEIVED = 2;
  this.LOADING = 3;
  this.DONE = 4;

  /**
   * Public vars
   */

  // Current state
  this.readyState = this.UNSENT;

  // default ready state change handler in case one is not set or is set late
  this.onreadystatechange = null;

  // Result & response
  this.responseText = "";
  this.responseXML = "";
  this.status = null;
  this.statusText = null;

  /**
   * Private methods
   */

  /**
   * Check if the specified header is allowed.
   *
   * @param string header Header to validate
   * @return boolean False if not allowed, otherwise true
   */
  var isAllowedHttpHeader = function isAllowedHttpHeader(header) {
    return disableHeaderCheck || header && forbiddenRequestHeaders.indexOf(header.toLowerCase()) === -1;
  };

  /**
   * Check if the specified method is allowed.
   *
   * @param string method Request method to validate
   * @return boolean False if not allowed, otherwise true
   */
  var isAllowedHttpMethod = function isAllowedHttpMethod(method) {
    return method && forbiddenRequestMethods.indexOf(method) === -1;
  };

  /**
   * Public methods
   */

  /**
   * Open the connection. Currently supports local server requests.
   *
   * @param string method Connection method (eg GET, POST)
   * @param string url URL for the connection.
   * @param boolean async Asynchronous connection. Default is true.
   * @param string user Username for basic authentication (optional)
   * @param string password Password for basic authentication (optional)
   */
  this.open = function (method, url, async, user, password) {
    this.abort();
    errorFlag = false;
    abortedFlag = false;

    // Check for valid request method
    if (!isAllowedHttpMethod(method)) {
      throw new Error("SecurityError: Request method not allowed");
    }
    settings = {
      "method": method,
      "url": url.toString(),
      "async": typeof async !== "boolean" ? true : async,
      "user": user || null,
      "password": password || null
    };
    setState(this.OPENED);
  };

  /**
   * Disables or enables isAllowedHttpHeader() check the request. Enabled by default.
   * This does not conform to the W3C spec.
   *
   * @param boolean state Enable or disable header checking.
   */
  this.setDisableHeaderCheck = function (state) {
    disableHeaderCheck = state;
  };

  /**
   * Sets a header for the request.
   *
   * @param string header Header name
   * @param string value Header value
   * @return boolean Header added
   */
  this.setRequestHeader = function (header, value) {
    if (this.readyState != this.OPENED) {
      throw new Error("INVALID_STATE_ERR: setRequestHeader can only be called when state is OPEN");
    }
    if (!isAllowedHttpHeader(header)) {
      console.warn('Refused to set unsafe header "' + header + '"');
      return false;
    }
    if (sendFlag) {
      throw new Error("INVALID_STATE_ERR: send flag is true");
    }
    headers[header] = value;
    return true;
  };

  /**
   * Gets a header from the server response.
   *
   * @param string header Name of header to get.
   * @return string Text of the header or null if it doesn't exist.
   */
  this.getResponseHeader = function (header) {
    if (typeof header === "string" && this.readyState > this.OPENED && response.headers[header.toLowerCase()] && !errorFlag) {
      return response.headers[header.toLowerCase()];
    }
    return null;
  };

  /**
   * Gets all the response headers.
   *
   * @return string A string with all response headers separated by CR+LF
   */
  this.getAllResponseHeaders = function () {
    if (this.readyState < this.HEADERS_RECEIVED || errorFlag) {
      return "";
    }
    var result = "";
    for (var i in response.headers) {
      // Cookie headers are excluded
      if (i !== "set-cookie" && i !== "set-cookie2") {
        result += i + ": " + response.headers[i] + "\r\n";
      }
    }
    return result.substr(0, result.length - 2);
  };

  /**
   * Gets a request header
   *
   * @param string name Name of header to get
   * @return string Returns the request header or empty string if not set
   */
  this.getRequestHeader = function (name) {
    // @TODO Make this case insensitive
    if (typeof name === "string" && headers[name]) {
      return headers[name];
    }
    return "";
  };

  /**
   * Sends the request to the server.
   *
   * @param string data Optional data to send as request body.
   */
  this.send = function (data) {
    if (this.readyState != this.OPENED) {
      throw new Error("INVALID_STATE_ERR: connection must be opened before send() is called");
    }
    if (sendFlag) {
      throw new Error("INVALID_STATE_ERR: send has already been called");
    }
    var ssl = false,
      local = false;
    var url = Url.parse(settings.url);
    var host;
    // Determine the server
    switch (url.protocol) {
      case 'https:':
        ssl = true;
      // SSL & non-SSL both need host, no break here.
      case 'http:':
        host = url.hostname;
        break;
      case 'file:':
        local = true;
        break;
      case undefined:
      case '':
        host = "localhost";
        break;
      default:
        throw new Error("Protocol not supported.");
    }

    // Load files off the local filesystem (file://)
    if (local) {
      if (settings.method !== "GET") {
        throw new Error("XMLHttpRequest: Only GET method is supported");
      }
      if (settings.async) {
        fs.readFile(unescape(url.pathname), 'utf8', function (error, data) {
          if (error) {
            self.handleError(error, error.errno || -1);
          } else {
            self.status = 200;
            self.responseText = data;
            setState(self.DONE);
          }
        });
      } else {
        try {
          this.responseText = fs.readFileSync(unescape(url.pathname), 'utf8');
          this.status = 200;
          setState(self.DONE);
        } catch (e) {
          this.handleError(e, e.errno || -1);
        }
      }
      return;
    }

    // Default to port 80. If accessing localhost on another port be sure
    // to use http://localhost:port/path
    var port = url.port || (ssl ? 443 : 80);
    // Add query string if one is used
    var uri = url.pathname + (url.search ? url.search : '');

    // Set the Host header or the server may reject the request
    headers["Host"] = host;
    if (!(ssl && port === 443 || port === 80)) {
      headers["Host"] += ':' + url.port;
    }

    // Set Basic Auth if necessary
    if (settings.user) {
      if (typeof settings.password == "undefined") {
        settings.password = "";
      }
      var authBuf = new Buffer(settings.user + ":" + settings.password);
      headers["Authorization"] = "Basic " + authBuf.toString("base64");
    }

    // Set content length header
    if (settings.method === "GET" || settings.method === "HEAD") {
      data = null;
    } else if (data) {
      headers["Content-Length"] = Buffer.isBuffer(data) ? data.length : Buffer.byteLength(data);
      if (!headers["Content-Type"]) {
        headers["Content-Type"] = "text/plain;charset=UTF-8";
      }
    } else if (settings.method === "POST") {
      // For a post with no data set Content-Length: 0.
      // This is required by buggy servers that don't meet the specs.
      headers["Content-Length"] = 0;
    }
    var agent = opts.agent || false;
    var options = {
      host: host,
      port: port,
      path: uri,
      method: settings.method,
      headers: headers,
      agent: agent
    };
    if (ssl) {
      options.pfx = opts.pfx;
      options.key = opts.key;
      options.passphrase = opts.passphrase;
      options.cert = opts.cert;
      options.ca = opts.ca;
      options.ciphers = opts.ciphers;
      options.rejectUnauthorized = opts.rejectUnauthorized === false ? false : true;
    }

    // Reset error flag
    errorFlag = false;
    // Handle async requests
    if (settings.async) {
      // Use the proper protocol
      var doRequest = ssl ? https.request : http.request;

      // Request is being sent, set send flag
      sendFlag = true;

      // As per spec, this is called here for historical reasons.
      self.dispatchEvent("readystatechange");

      // Handler for the response
      var _responseHandler = function responseHandler(resp) {
        // Set response var to the response we got back
        // This is so it remains accessable outside this scope
        response = resp;
        // Check for redirect
        // @TODO Prevent looped redirects
        if (response.statusCode === 302 || response.statusCode === 303 || response.statusCode === 307) {
          // Change URL to the redirect location
          settings.url = response.headers.location;
          var url = Url.parse(settings.url);
          // Set host var in case it's used later
          host = url.hostname;
          // Options for the new request
          var newOptions = {
            hostname: url.hostname,
            port: url.port,
            path: url.path,
            method: response.statusCode === 303 ? 'GET' : settings.method,
            headers: headers
          };
          if (ssl) {
            newOptions.pfx = opts.pfx;
            newOptions.key = opts.key;
            newOptions.passphrase = opts.passphrase;
            newOptions.cert = opts.cert;
            newOptions.ca = opts.ca;
            newOptions.ciphers = opts.ciphers;
            newOptions.rejectUnauthorized = opts.rejectUnauthorized === false ? false : true;
          }

          // Issue the new request
          request = doRequest(newOptions, _responseHandler).on('error', errorHandler);
          request.end();
          // @TODO Check if an XHR event needs to be fired here
          return;
        }
        if (response && response.setEncoding) {
          response.setEncoding("utf8");
        }
        setState(self.HEADERS_RECEIVED);
        self.status = response.statusCode;
        response.on('data', function (chunk) {
          // Make sure there's some data
          if (chunk) {
            self.responseText += chunk;
          }
          // Don't emit state changes if the connection has been aborted.
          if (sendFlag) {
            setState(self.LOADING);
          }
        });
        response.on('end', function () {
          if (sendFlag) {
            // The sendFlag needs to be set before setState is called.  Otherwise if we are chaining callbacks
            // there can be a timing issue (the callback is called and a new call is made before the flag is reset).
            sendFlag = false;
            // Discard the 'end' event if the connection has been aborted
            setState(self.DONE);
          }
        });
        response.on('error', function (error) {
          self.handleError(error);
        });
      };

      // Error handler for the request
      var errorHandler = function errorHandler(error) {
        self.handleError(error);
      };

      // Create the request
      request = doRequest(options, _responseHandler).on('error', errorHandler);
      if (opts.autoUnref) {
        request.on('socket', function (socket) {
          socket.unref();
        });
      }

      // Node 0.4 and later won't accept empty data. Make sure it's needed.
      if (data) {
        request.write(data);
      }
      request.end();
      self.dispatchEvent("loadstart");
    } else {
      // Synchronous
      // Create a temporary file for communication with the other Node process
      var contentFile = ".node-xmlhttprequest-content-" + process.pid;
      var syncFile = ".node-xmlhttprequest-sync-" + process.pid;
      fs.writeFileSync(syncFile, "", "utf8");
      // The async request the other Node process executes
      var execString = "var http = require('http'), https = require('https'), fs = require('fs');" + "var doRequest = http" + (ssl ? "s" : "") + ".request;" + "var options = " + JSON.stringify(options) + ";" + "var responseText = '';" + "var req = doRequest(options, function(response) {" + "response.setEncoding('utf8');" + "response.on('data', function(chunk) {" + "  responseText += chunk;" + "});" + "response.on('end', function() {" + "fs.writeFileSync('" + contentFile + "', 'NODE-XMLHTTPREQUEST-STATUS:' + response.statusCode + ',' + responseText, 'utf8');" + "fs.unlinkSync('" + syncFile + "');" + "});" + "response.on('error', function(error) {" + "fs.writeFileSync('" + contentFile + "', 'NODE-XMLHTTPREQUEST-ERROR:' + JSON.stringify(error), 'utf8');" + "fs.unlinkSync('" + syncFile + "');" + "});" + "}).on('error', function(error) {" + "fs.writeFileSync('" + contentFile + "', 'NODE-XMLHTTPREQUEST-ERROR:' + JSON.stringify(error), 'utf8');" + "fs.unlinkSync('" + syncFile + "');" + "});" + (data ? "req.write('" + JSON.stringify(data).slice(1, -1).replace(/'/g, "\\'") + "');" : "") + "req.end();";
      // Start the other Node Process, executing this string
      var syncProc = spawn(process.argv[0], ["-e", execString]);
      while (fs.existsSync(syncFile)) {
        // Wait while the sync file is empty
      }
      self.responseText = fs.readFileSync(contentFile, 'utf8');
      // Kill the child process once the file has data
      syncProc.stdin.end();
      // Remove the temporary file
      fs.unlinkSync(contentFile);
      if (self.responseText.match(/^NODE-XMLHTTPREQUEST-ERROR:/)) {
        // If the file returned an error, handle it
        var errorObj = self.responseText.replace(/^NODE-XMLHTTPREQUEST-ERROR:/, "");
        self.handleError(errorObj, 503);
      } else {
        // If the file returned okay, parse its data and move to the DONE state
        self.status = self.responseText.replace(/^NODE-XMLHTTPREQUEST-STATUS:([0-9]*),.*/, "$1");
        self.responseText = self.responseText.replace(/^NODE-XMLHTTPREQUEST-STATUS:[0-9]*,(.*)/, "$1");
        setState(self.DONE);
      }
    }
  };

  /**
   * Called when an error is encountered to deal with it.
   * @param  status  {number}    HTTP status code to use rather than the default (0) for XHR errors.
   */
  this.handleError = function (error, status) {
    this.status = status || 0;
    this.statusText = error;
    this.responseText = error.stack;
    errorFlag = true;
    setState(this.DONE);
  };

  /**
   * Aborts a request.
   */
  this.abort = function () {
    if (request) {
      request.abort();
      request = null;
    }
    headers = Object.assign({}, defaultHeaders);
    this.responseText = "";
    this.responseXML = "";
    errorFlag = abortedFlag = true;
    if (this.readyState !== this.UNSENT && (this.readyState !== this.OPENED || sendFlag) && this.readyState !== this.DONE) {
      sendFlag = false;
      setState(this.DONE);
    }
    this.readyState = this.UNSENT;
  };

  /**
   * Adds an event listener. Preferred method of binding to events.
   */
  this.addEventListener = function (event, callback) {
    if (!(event in listeners)) {
      listeners[event] = [];
    }
    // Currently allows duplicate callbacks. Should it?
    listeners[event].push(callback);
  };

  /**
   * Remove an event callback that has already been bound.
   * Only works on the matching funciton, cannot be a copy.
   */
  this.removeEventListener = function (event, callback) {
    if (event in listeners) {
      // Filter will return a new array with the callback removed
      listeners[event] = listeners[event].filter(function (ev) {
        return ev !== callback;
      });
    }
  };

  /**
   * Dispatch any events, including both "on" methods and events attached using addEventListener.
   */
  this.dispatchEvent = function (event) {
    var _this = this;
    if (typeof self["on" + event] === "function") {
      if (this.readyState === this.DONE) setImmediate(function () {
        self["on" + event]();
      });else self["on" + event]();
    }
    if (event in listeners) {
      var _loop = function _loop(i) {
        if (_this.readyState === _this.DONE) setImmediate(function () {
          listeners[event][i].call(self);
        });else listeners[event][i].call(self);
      };
      for (var i = 0, len = listeners[event].length; i < len; i++) {
        _loop(i);
      }
    }
  };

  /**
   * Changes readyState and calls onreadystatechange.
   *
   * @param int state New state
   */
  var setState = function setState(state) {
    if (self.readyState === state || self.readyState === self.UNSENT && abortedFlag) return;
    self.readyState = state;
    if (settings.async || self.readyState < self.OPENED || self.readyState === self.DONE) {
      self.dispatchEvent("readystatechange");
    }
    if (self.readyState === self.DONE) {
      var fire;
      if (abortedFlag) fire = "abort";else if (errorFlag) fire = "error";else fire = "load";
      self.dispatchEvent(fire);

      // @TODO figure out InspectorInstrumentation::didLoadXHR(cookie)
      self.dispatchEvent("loadend");
    }
  };
}
var XMLHttpRequest$1 = /*@__PURE__*/getDefaultExportFromCjs(XMLHttpRequest_1);

var XMLHttpRequestModule = /*#__PURE__*/_mergeNamespaces({
    __proto__: null,
    default: XMLHttpRequest$1
}, [XMLHttpRequest_1]);

var XHR = XMLHttpRequest$1 || XMLHttpRequestModule;
function createCookieJar() {
  return new CookieJar();
}
/**
 * @see https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie
 */
function parse$3(setCookieString) {
  var parts = setCookieString.split("; ");
  var i = parts[0].indexOf("=");
  if (i === -1) {
    return;
  }
  var name = parts[0].substring(0, i).trim();
  if (!name.length) {
    return;
  }
  var value = parts[0].substring(i + 1).trim();
  if (value.charCodeAt(0) === 0x22) {
    // remove double quotes
    value = value.slice(1, -1);
  }
  var cookie = {
    name: name,
    value: value
  };
  for (var j = 1; j < parts.length; j++) {
    var subParts = parts[j].split("=");
    if (subParts.length !== 2) {
      continue;
    }
    var key = subParts[0].trim();
    var _value = subParts[1].trim();
    switch (key) {
      case "Expires":
        cookie.expires = new Date(_value);
        break;
      case "Max-Age":
        var expiration = new Date();
        expiration.setUTCSeconds(expiration.getUTCSeconds() + parseInt(_value, 10));
        cookie.expires = expiration;
        break;
      // ignore other keys
    }
  }
  return cookie;
}
var CookieJar = /*#__PURE__*/function () {
  function CookieJar() {
    _classCallCheck(this, CookieJar);
    this.cookies = new Map();
  }
  return _createClass(CookieJar, [{
    key: "parseCookies",
    value: function parseCookies(xhr) {
      var _this = this;
      var values = xhr.getResponseHeader("set-cookie");
      if (!values) {
        return;
      }
      values.forEach(function (value) {
        var parsed = parse$3(value);
        if (parsed) {
          _this.cookies.set(parsed.name, parsed);
        }
      });
    }
  }, {
    key: "addCookies",
    value: function addCookies(xhr) {
      var _this2 = this;
      var cookies = [];
      this.cookies.forEach(function (cookie, name) {
        var _a;
        if (((_a = cookie.expires) === null || _a === void 0 ? void 0 : _a.getTime()) < Date.now()) {
          _this2.cookies["delete"](name);
        } else {
          cookies.push("".concat(name, "=").concat(cookie.value));
        }
      });
      if (cookies.length) {
        xhr.setDisableHeaderCheck(true);
        xhr.setRequestHeader("cookie", cookies.join("; "));
      }
    }
  }]);
}();

function empty() {}
var hasXHR2 = function () {
  var xhr = new XHR({
    xdomain: false
  });
  return null != xhr.responseType;
}();
var Polling = /*#__PURE__*/function (_Transport) {
  /**
   * XHR Polling constructor.
   *
   * @param {Object} opts
   * @package
   */
  function Polling(opts) {
    var _this;
    _classCallCheck(this, Polling);
    _this = _callSuper(this, Polling, [opts]);
    _this.polling = false;
    if (typeof location !== "undefined") {
      var isSSL = "https:" === location.protocol;
      var port = location.port;
      // some user agents have empty `location.port`
      if (!port) {
        port = isSSL ? "443" : "80";
      }
      _this.xd = typeof location !== "undefined" && opts.hostname !== location.hostname || port !== opts.port;
    }
    /**
     * XHR supports binary
     */
    var forceBase64 = opts && opts.forceBase64;
    _this.supportsBinary = hasXHR2 && !forceBase64;
    if (_this.opts.withCredentials) {
      _this.cookieJar = createCookieJar();
    }
    return _this;
  }
  _inherits(Polling, _Transport);
  return _createClass(Polling, [{
    key: "name",
    get: function get() {
      return "polling";
    }
    /**
     * Opens the socket (triggers polling). We write a PING message to determine
     * when the transport is open.
     *
     * @protected
     */
  }, {
    key: "doOpen",
    value: function doOpen() {
      this.poll();
    }
    /**
     * Pauses polling.
     *
     * @param {Function} onPause - callback upon buffers are flushed and transport is paused
     * @package
     */
  }, {
    key: "pause",
    value: function pause(onPause) {
      var _this2 = this;
      this.readyState = "pausing";
      var pause = function pause() {
        _this2.readyState = "paused";
        onPause();
      };
      if (this.polling || !this.writable) {
        var total = 0;
        if (this.polling) {
          total++;
          this.once("pollComplete", function () {
            --total || pause();
          });
        }
        if (!this.writable) {
          total++;
          this.once("drain", function () {
            --total || pause();
          });
        }
      } else {
        pause();
      }
    }
    /**
     * Starts polling cycle.
     *
     * @private
     */
  }, {
    key: "poll",
    value: function poll() {
      this.polling = true;
      this.doPoll();
      this.emitReserved("poll");
    }
    /**
     * Overloads onData to detect payloads.
     *
     * @protected
     */
  }, {
    key: "onData",
    value: function onData(data) {
      var _this3 = this;
      var callback = function callback(packet) {
        // if its the first message we consider the transport open
        if ("opening" === _this3.readyState && packet.type === "open") {
          _this3.onOpen();
        }
        // if its a close packet, we close the ongoing requests
        if ("close" === packet.type) {
          _this3.onClose({
            description: "transport closed by the server"
          });
          return false;
        }
        // otherwise bypass onData and handle the message
        _this3.onPacket(packet);
      };
      // decode payload
      decodePayload(data, this.socket.binaryType).forEach(callback);
      // if an event did not trigger closing
      if ("closed" !== this.readyState) {
        // if we got data we're not polling
        this.polling = false;
        this.emitReserved("pollComplete");
        if ("open" === this.readyState) {
          this.poll();
        }
      }
    }
    /**
     * For polling, send a close packet.
     *
     * @protected
     */
  }, {
    key: "doClose",
    value: function doClose() {
      var _this4 = this;
      var close = function close() {
        _this4.write([{
          type: "close"
        }]);
      };
      if ("open" === this.readyState) {
        close();
      } else {
        // in case we're trying to close while
        // handshaking is in progress (GH-164)
        this.once("open", close);
      }
    }
    /**
     * Writes a packets payload.
     *
     * @param {Array} packets - data packets
     * @protected
     */
  }, {
    key: "write",
    value: function write(packets) {
      var _this5 = this;
      this.writable = false;
      encodePayload(packets, function (data) {
        _this5.doWrite(data, function () {
          _this5.writable = true;
          _this5.emitReserved("drain");
        });
      });
    }
    /**
     * Generates uri for connection.
     *
     * @private
     */
  }, {
    key: "uri",
    value: function uri() {
      var schema = this.opts.secure ? "https" : "http";
      var query = this.query || {};
      // cache busting is forced
      if (false !== this.opts.timestampRequests) {
        query[this.opts.timestampParam] = yeast();
      }
      if (!this.supportsBinary && !query.sid) {
        query.b64 = 1;
      }
      return this.createUri(schema, query);
    }
    /**
     * Creates a request.
     *
     * @param {String} method
     * @private
     */
  }, {
    key: "request",
    value: function request() {
      var opts = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {};
      Object.assign(opts, {
        xd: this.xd,
        cookieJar: this.cookieJar
      }, this.opts);
      return new Request(this.uri(), opts);
    }
    /**
     * Sends data.
     *
     * @param {String} data to send.
     * @param {Function} called upon flush.
     * @private
     */
  }, {
    key: "doWrite",
    value: function doWrite(data, fn) {
      var _this6 = this;
      var req = this.request({
        method: "POST",
        data: data
      });
      req.on("success", fn);
      req.on("error", function (xhrStatus, context) {
        _this6.onError("xhr post error", xhrStatus, context);
      });
    }
    /**
     * Starts a poll cycle.
     *
     * @private
     */
  }, {
    key: "doPoll",
    value: function doPoll() {
      var _this7 = this;
      var req = this.request();
      req.on("data", this.onData.bind(this));
      req.on("error", function (xhrStatus, context) {
        _this7.onError("xhr poll error", xhrStatus, context);
      });
      this.pollXhr = req;
    }
  }]);
}(Transport);
var Request = /*#__PURE__*/function (_Emitter) {
  /**
   * Request constructor
   *
   * @param {Object} options
   * @package
   */
  function Request(uri, opts) {
    var _this8;
    _classCallCheck(this, Request);
    _this8 = _callSuper(this, Request);
    installTimerFunctions(_this8, opts);
    _this8.opts = opts;
    _this8.method = opts.method || "GET";
    _this8.uri = uri;
    _this8.data = undefined !== opts.data ? opts.data : null;
    _this8.create();
    return _this8;
  }
  /**
   * Creates the XHR object and sends the request.
   *
   * @private
   */
  _inherits(Request, _Emitter);
  return _createClass(Request, [{
    key: "create",
    value: function create() {
      var _this9 = this;
      var _a;
      var opts = pick(this.opts, "agent", "pfx", "key", "passphrase", "cert", "ca", "ciphers", "rejectUnauthorized", "autoUnref");
      opts.xdomain = !!this.opts.xd;
      var xhr = this.xhr = new XHR(opts);
      try {
        xhr.open(this.method, this.uri, true);
        try {
          if (this.opts.extraHeaders) {
            xhr.setDisableHeaderCheck && xhr.setDisableHeaderCheck(true);
            for (var i in this.opts.extraHeaders) {
              if (this.opts.extraHeaders.hasOwnProperty(i)) {
                xhr.setRequestHeader(i, this.opts.extraHeaders[i]);
              }
            }
          }
        } catch (e) {}
        if ("POST" === this.method) {
          try {
            xhr.setRequestHeader("Content-type", "text/plain;charset=UTF-8");
          } catch (e) {}
        }
        try {
          xhr.setRequestHeader("Accept", "*/*");
        } catch (e) {}
        (_a = this.opts.cookieJar) === null || _a === void 0 ? void 0 : _a.addCookies(xhr);
        // ie6 check
        if ("withCredentials" in xhr) {
          xhr.withCredentials = this.opts.withCredentials;
        }
        if (this.opts.requestTimeout) {
          xhr.timeout = this.opts.requestTimeout;
        }
        xhr.onreadystatechange = function () {
          var _a;
          if (xhr.readyState === 3) {
            (_a = _this9.opts.cookieJar) === null || _a === void 0 ? void 0 : _a.parseCookies(xhr);
          }
          if (4 !== xhr.readyState) return;
          if (200 === xhr.status || 1223 === xhr.status) {
            _this9.onLoad();
          } else {
            // make sure the `error` event handler that's user-set
            // does not throw in the same tick and gets caught here
            _this9.setTimeoutFn(function () {
              _this9.onError(typeof xhr.status === "number" ? xhr.status : 0);
            }, 0);
          }
        };
        xhr.send(this.data);
      } catch (e) {
        // Need to defer since .create() is called directly from the constructor
        // and thus the 'error' event can only be only bound *after* this exception
        // occurs.  Therefore, also, we cannot throw here at all.
        this.setTimeoutFn(function () {
          _this9.onError(e);
        }, 0);
        return;
      }
      if (typeof document !== "undefined") {
        this.index = Request.requestsCount++;
        Request.requests[this.index] = this;
      }
    }
    /**
     * Called upon error.
     *
     * @private
     */
  }, {
    key: "onError",
    value: function onError(err) {
      this.emitReserved("error", err, this.xhr);
      this.cleanup(true);
    }
    /**
     * Cleans up house.
     *
     * @private
     */
  }, {
    key: "cleanup",
    value: function cleanup(fromError) {
      if ("undefined" === typeof this.xhr || null === this.xhr) {
        return;
      }
      this.xhr.onreadystatechange = empty;
      if (fromError) {
        try {
          this.xhr.abort();
        } catch (e) {}
      }
      if (typeof document !== "undefined") {
        delete Request.requests[this.index];
      }
      this.xhr = null;
    }
    /**
     * Called upon load.
     *
     * @private
     */
  }, {
    key: "onLoad",
    value: function onLoad() {
      var data = this.xhr.responseText;
      if (data !== null) {
        this.emitReserved("data", data);
        this.emitReserved("success");
        this.cleanup();
      }
    }
    /**
     * Aborts the request.
     *
     * @package
     */
  }, {
    key: "abort",
    value: function abort() {
      this.cleanup();
    }
  }]);
}(Emitter);
Request.requestsCount = 0;
Request.requests = {};
/**
 * Aborts pending requests when unloading the window. This is needed to prevent
 * memory leaks (e.g. when using IE) and to ensure that no spurious error is
 * emitted.
 */
if (typeof document !== "undefined") {
  // @ts-ignore
  if (typeof attachEvent === "function") {
    // @ts-ignore
    attachEvent("onunload", unloadHandler);
  } else if (typeof addEventListener === "function") {
    var terminationEvent = "onpagehide" in globalThisShim ? "pagehide" : "unload";
    addEventListener(terminationEvent, unloadHandler, false);
  }
}
function unloadHandler() {
  for (var i in Request.requests) {
    if (Request.requests.hasOwnProperty(i)) {
      Request.requests[i].abort();
    }
  }
}

require$$0$1.Duplex;

var bufferUtil$2 = {exports: {}};

var constants = {
  BINARY_TYPES: ['nodebuffer', 'arraybuffer', 'fragments'],
  EMPTY_BUFFER: Buffer.alloc(0),
  GUID: '258EAFA5-E914-47DA-95CA-C5AB0DC85B11',
  kForOnEventAttribute: Symbol('kIsForOnEventAttribute'),
  kListener: Symbol('kListener'),
  kStatusCode: Symbol('status-code'),
  kWebSocket: Symbol('websocket'),
  NOOP: function NOOP() {}
};

var unmask$1;
var mask;
var EMPTY_BUFFER$3 = constants.EMPTY_BUFFER;
var FastBuffer$2 = Buffer[Symbol.species];

/**
 * Merges an array of buffers into a new buffer.
 *
 * @param {Buffer[]} list The array of buffers to concat
 * @param {Number} totalLength The total length of buffers in the list
 * @return {Buffer} The resulting buffer
 * @public
 */
function concat$1(list, totalLength) {
  if (list.length === 0) return EMPTY_BUFFER$3;
  if (list.length === 1) return list[0];
  var target = Buffer.allocUnsafe(totalLength);
  var offset = 0;
  for (var i = 0; i < list.length; i++) {
    var buf = list[i];
    target.set(buf, offset);
    offset += buf.length;
  }
  if (offset < totalLength) {
    return new FastBuffer$2(target.buffer, target.byteOffset, offset);
  }
  return target;
}

/**
 * Masks a buffer using the given mask.
 *
 * @param {Buffer} source The buffer to mask
 * @param {Buffer} mask The mask to use
 * @param {Buffer} output The buffer where to store the result
 * @param {Number} offset The offset at which to start writing
 * @param {Number} length The number of bytes to mask.
 * @public
 */
function _mask(source, mask, output, offset, length) {
  for (var i = 0; i < length; i++) {
    output[offset + i] = source[i] ^ mask[i & 3];
  }
}

/**
 * Unmasks a buffer using the given mask.
 *
 * @param {Buffer} buffer The buffer to unmask
 * @param {Buffer} mask The mask to use
 * @public
 */
function _unmask(buffer, mask) {
  for (var i = 0; i < buffer.length; i++) {
    buffer[i] ^= mask[i & 3];
  }
}

/**
 * Converts a buffer to an `ArrayBuffer`.
 *
 * @param {Buffer} buf The buffer to convert
 * @return {ArrayBuffer} Converted buffer
 * @public
 */
function toArrayBuffer$1(buf) {
  if (buf.length === buf.buffer.byteLength) {
    return buf.buffer;
  }
  return buf.buffer.slice(buf.byteOffset, buf.byteOffset + buf.length);
}

/**
 * Converts `data` to a `Buffer`.
 *
 * @param {*} data The data to convert
 * @return {Buffer} The buffer
 * @throws {TypeError}
 * @public
 */
function toBuffer$2(data) {
  toBuffer$2.readOnly = true;
  if (Buffer.isBuffer(data)) return data;
  var buf;
  if (data instanceof ArrayBuffer) {
    buf = new FastBuffer$2(data);
  } else if (ArrayBuffer.isView(data)) {
    buf = new FastBuffer$2(data.buffer, data.byteOffset, data.byteLength);
  } else {
    buf = Buffer.from(data);
    toBuffer$2.readOnly = false;
  }
  return buf;
}
bufferUtil$2.exports = {
  concat: concat$1,
  mask: _mask,
  toArrayBuffer: toArrayBuffer$1,
  toBuffer: toBuffer$2,
  unmask: _unmask
};

/* istanbul ignore else  */
if (!process.env.WS_NO_BUFFER_UTIL) {
  try {
    var bufferUtil$1 = require('bufferutil');
    mask = bufferUtil$2.exports.mask = function (source, mask, output, offset, length) {
      if (length < 48) _mask(source, mask, output, offset, length);else bufferUtil$1.mask(source, mask, output, offset, length);
    };
    unmask$1 = bufferUtil$2.exports.unmask = function (buffer, mask) {
      if (buffer.length < 32) _unmask(buffer, mask);else bufferUtil$1.unmask(buffer, mask);
    };
  } catch (e) {
    // Continue regardless of the error.
  }
}
var bufferUtilExports = bufferUtil$2.exports;

var kDone = Symbol('kDone');
var kRun = Symbol('kRun');

/**
 * A very simple job queue with adjustable concurrency. Adapted from
 * https://github.com/STRML/async-limiter
 */
var Limiter$1 = /*#__PURE__*/function () {
  /**
   * Creates a new `Limiter`.
   *
   * @param {Number} [concurrency=Infinity] The maximum number of jobs allowed
   *     to run concurrently
   */
  function Limiter(concurrency) {
    var _this = this;
    _classCallCheck(this, Limiter);
    this[kDone] = function () {
      _this.pending--;
      _this[kRun]();
    };
    this.concurrency = concurrency || Infinity;
    this.jobs = [];
    this.pending = 0;
  }

  /**
   * Adds a job to the queue.
   *
   * @param {Function} job The job to run
   * @public
   */
  return _createClass(Limiter, [{
    key: "add",
    value: function add(job) {
      this.jobs.push(job);
      this[kRun]();
    }

    /**
     * Removes a job from the queue and runs it if possible.
     *
     * @private
     */
  }, {
    key: kRun,
    value: function value() {
      if (this.pending === this.concurrency) return;
      if (this.jobs.length) {
        var job = this.jobs.shift();
        this.pending++;
        job(this[kDone]);
      }
    }
  }]);
}();
var limiter = Limiter$1;

var zlib = require$$0$2;
var bufferUtil = bufferUtilExports;
var Limiter = limiter;
var kStatusCode$2 = constants.kStatusCode;
var FastBuffer$1 = Buffer[Symbol.species];
var TRAILER = Buffer.from([0x00, 0x00, 0xff, 0xff]);
var kPerMessageDeflate = Symbol('permessage-deflate');
var kTotalLength = Symbol('total-length');
var kCallback = Symbol('callback');
var kBuffers = Symbol('buffers');
var kError$1 = Symbol('error');

//
// We limit zlib concurrency, which prevents severe memory fragmentation
// as documented in https://github.com/nodejs/node/issues/8871#issuecomment-250915913
// and https://github.com/websockets/ws/issues/1202
//
// Intentionally global; it's the global thread pool that's an issue.
//
var zlibLimiter;

/**
 * permessage-deflate implementation.
 */
var PerMessageDeflate$3 = /*#__PURE__*/function () {
  /**
   * Creates a PerMessageDeflate instance.
   *
   * @param {Object} [options] Configuration options
   * @param {(Boolean|Number)} [options.clientMaxWindowBits] Advertise support
   *     for, or request, a custom client window size
   * @param {Boolean} [options.clientNoContextTakeover=false] Advertise/
   *     acknowledge disabling of client context takeover
   * @param {Number} [options.concurrencyLimit=10] The number of concurrent
   *     calls to zlib
   * @param {(Boolean|Number)} [options.serverMaxWindowBits] Request/confirm the
   *     use of a custom server window size
   * @param {Boolean} [options.serverNoContextTakeover=false] Request/accept
   *     disabling of server context takeover
   * @param {Number} [options.threshold=1024] Size (in bytes) below which
   *     messages should not be compressed if context takeover is disabled
   * @param {Object} [options.zlibDeflateOptions] Options to pass to zlib on
   *     deflate
   * @param {Object} [options.zlibInflateOptions] Options to pass to zlib on
   *     inflate
   * @param {Boolean} [isServer=false] Create the instance in either server or
   *     client mode
   * @param {Number} [maxPayload=0] The maximum allowed message length
   */
  function PerMessageDeflate(options, isServer, maxPayload) {
    _classCallCheck(this, PerMessageDeflate);
    this._maxPayload = maxPayload | 0;
    this._options = options || {};
    this._threshold = this._options.threshold !== undefined ? this._options.threshold : 1024;
    this._isServer = !!isServer;
    this._deflate = null;
    this._inflate = null;
    this.params = null;
    if (!zlibLimiter) {
      var concurrency = this._options.concurrencyLimit !== undefined ? this._options.concurrencyLimit : 10;
      zlibLimiter = new Limiter(concurrency);
    }
  }

  /**
   * @type {String}
   */
  return _createClass(PerMessageDeflate, [{
    key: "offer",
    value:
    /**
     * Create an extension negotiation offer.
     *
     * @return {Object} Extension parameters
     * @public
     */
    function offer() {
      var params = {};
      if (this._options.serverNoContextTakeover) {
        params.server_no_context_takeover = true;
      }
      if (this._options.clientNoContextTakeover) {
        params.client_no_context_takeover = true;
      }
      if (this._options.serverMaxWindowBits) {
        params.server_max_window_bits = this._options.serverMaxWindowBits;
      }
      if (this._options.clientMaxWindowBits) {
        params.client_max_window_bits = this._options.clientMaxWindowBits;
      } else if (this._options.clientMaxWindowBits == null) {
        params.client_max_window_bits = true;
      }
      return params;
    }

    /**
     * Accept an extension negotiation offer/response.
     *
     * @param {Array} configurations The extension negotiation offers/reponse
     * @return {Object} Accepted configuration
     * @public
     */
  }, {
    key: "accept",
    value: function accept(configurations) {
      configurations = this.normalizeParams(configurations);
      this.params = this._isServer ? this.acceptAsServer(configurations) : this.acceptAsClient(configurations);
      return this.params;
    }

    /**
     * Releases all resources used by the extension.
     *
     * @public
     */
  }, {
    key: "cleanup",
    value: function cleanup() {
      if (this._inflate) {
        this._inflate.close();
        this._inflate = null;
      }
      if (this._deflate) {
        var callback = this._deflate[kCallback];
        this._deflate.close();
        this._deflate = null;
        if (callback) {
          callback(new Error('The deflate stream was closed while data was being processed'));
        }
      }
    }

    /**
     *  Accept an extension negotiation offer.
     *
     * @param {Array} offers The extension negotiation offers
     * @return {Object} Accepted configuration
     * @private
     */
  }, {
    key: "acceptAsServer",
    value: function acceptAsServer(offers) {
      var opts = this._options;
      var accepted = offers.find(function (params) {
        if (opts.serverNoContextTakeover === false && params.server_no_context_takeover || params.server_max_window_bits && (opts.serverMaxWindowBits === false || typeof opts.serverMaxWindowBits === 'number' && opts.serverMaxWindowBits > params.server_max_window_bits) || typeof opts.clientMaxWindowBits === 'number' && !params.client_max_window_bits) {
          return false;
        }
        return true;
      });
      if (!accepted) {
        throw new Error('None of the extension offers can be accepted');
      }
      if (opts.serverNoContextTakeover) {
        accepted.server_no_context_takeover = true;
      }
      if (opts.clientNoContextTakeover) {
        accepted.client_no_context_takeover = true;
      }
      if (typeof opts.serverMaxWindowBits === 'number') {
        accepted.server_max_window_bits = opts.serverMaxWindowBits;
      }
      if (typeof opts.clientMaxWindowBits === 'number') {
        accepted.client_max_window_bits = opts.clientMaxWindowBits;
      } else if (accepted.client_max_window_bits === true || opts.clientMaxWindowBits === false) {
        delete accepted.client_max_window_bits;
      }
      return accepted;
    }

    /**
     * Accept the extension negotiation response.
     *
     * @param {Array} response The extension negotiation response
     * @return {Object} Accepted configuration
     * @private
     */
  }, {
    key: "acceptAsClient",
    value: function acceptAsClient(response) {
      var params = response[0];
      if (this._options.clientNoContextTakeover === false && params.client_no_context_takeover) {
        throw new Error('Unexpected parameter "client_no_context_takeover"');
      }
      if (!params.client_max_window_bits) {
        if (typeof this._options.clientMaxWindowBits === 'number') {
          params.client_max_window_bits = this._options.clientMaxWindowBits;
        }
      } else if (this._options.clientMaxWindowBits === false || typeof this._options.clientMaxWindowBits === 'number' && params.client_max_window_bits > this._options.clientMaxWindowBits) {
        throw new Error('Unexpected or invalid parameter "client_max_window_bits"');
      }
      return params;
    }

    /**
     * Normalize parameters.
     *
     * @param {Array} configurations The extension negotiation offers/reponse
     * @return {Array} The offers/response with normalized parameters
     * @private
     */
  }, {
    key: "normalizeParams",
    value: function normalizeParams(configurations) {
      var _this = this;
      configurations.forEach(function (params) {
        Object.keys(params).forEach(function (key) {
          var value = params[key];
          if (value.length > 1) {
            throw new Error("Parameter \"".concat(key, "\" must have only a single value"));
          }
          value = value[0];
          if (key === 'client_max_window_bits') {
            if (value !== true) {
              var num = +value;
              if (!Number.isInteger(num) || num < 8 || num > 15) {
                throw new TypeError("Invalid value for parameter \"".concat(key, "\": ").concat(value));
              }
              value = num;
            } else if (!_this._isServer) {
              throw new TypeError("Invalid value for parameter \"".concat(key, "\": ").concat(value));
            }
          } else if (key === 'server_max_window_bits') {
            var _num = +value;
            if (!Number.isInteger(_num) || _num < 8 || _num > 15) {
              throw new TypeError("Invalid value for parameter \"".concat(key, "\": ").concat(value));
            }
            value = _num;
          } else if (key === 'client_no_context_takeover' || key === 'server_no_context_takeover') {
            if (value !== true) {
              throw new TypeError("Invalid value for parameter \"".concat(key, "\": ").concat(value));
            }
          } else {
            throw new Error("Unknown parameter \"".concat(key, "\""));
          }
          params[key] = value;
        });
      });
      return configurations;
    }

    /**
     * Decompress data. Concurrency limited.
     *
     * @param {Buffer} data Compressed data
     * @param {Boolean} fin Specifies whether or not this is the last fragment
     * @param {Function} callback Callback
     * @public
     */
  }, {
    key: "decompress",
    value: function decompress(data, fin, callback) {
      var _this2 = this;
      zlibLimiter.add(function (done) {
        _this2._decompress(data, fin, function (err, result) {
          done();
          callback(err, result);
        });
      });
    }

    /**
     * Compress data. Concurrency limited.
     *
     * @param {(Buffer|String)} data Data to compress
     * @param {Boolean} fin Specifies whether or not this is the last fragment
     * @param {Function} callback Callback
     * @public
     */
  }, {
    key: "compress",
    value: function compress(data, fin, callback) {
      var _this3 = this;
      zlibLimiter.add(function (done) {
        _this3._compress(data, fin, function (err, result) {
          done();
          callback(err, result);
        });
      });
    }

    /**
     * Decompress data.
     *
     * @param {Buffer} data Compressed data
     * @param {Boolean} fin Specifies whether or not this is the last fragment
     * @param {Function} callback Callback
     * @private
     */
  }, {
    key: "_decompress",
    value: function _decompress(data, fin, callback) {
      var _this4 = this;
      var endpoint = this._isServer ? 'client' : 'server';
      if (!this._inflate) {
        var key = "".concat(endpoint, "_max_window_bits");
        var windowBits = typeof this.params[key] !== 'number' ? zlib.Z_DEFAULT_WINDOWBITS : this.params[key];
        this._inflate = zlib.createInflateRaw(_objectSpread2(_objectSpread2({}, this._options.zlibInflateOptions), {}, {
          windowBits: windowBits
        }));
        this._inflate[kPerMessageDeflate] = this;
        this._inflate[kTotalLength] = 0;
        this._inflate[kBuffers] = [];
        this._inflate.on('error', inflateOnError);
        this._inflate.on('data', inflateOnData);
      }
      this._inflate[kCallback] = callback;
      this._inflate.write(data);
      if (fin) this._inflate.write(TRAILER);
      this._inflate.flush(function () {
        var err = _this4._inflate[kError$1];
        if (err) {
          _this4._inflate.close();
          _this4._inflate = null;
          callback(err);
          return;
        }
        var data = bufferUtil.concat(_this4._inflate[kBuffers], _this4._inflate[kTotalLength]);
        if (_this4._inflate._readableState.endEmitted) {
          _this4._inflate.close();
          _this4._inflate = null;
        } else {
          _this4._inflate[kTotalLength] = 0;
          _this4._inflate[kBuffers] = [];
          if (fin && _this4.params["".concat(endpoint, "_no_context_takeover")]) {
            _this4._inflate.reset();
          }
        }
        callback(null, data);
      });
    }

    /**
     * Compress data.
     *
     * @param {(Buffer|String)} data Data to compress
     * @param {Boolean} fin Specifies whether or not this is the last fragment
     * @param {Function} callback Callback
     * @private
     */
  }, {
    key: "_compress",
    value: function _compress(data, fin, callback) {
      var _this5 = this;
      var endpoint = this._isServer ? 'server' : 'client';
      if (!this._deflate) {
        var key = "".concat(endpoint, "_max_window_bits");
        var windowBits = typeof this.params[key] !== 'number' ? zlib.Z_DEFAULT_WINDOWBITS : this.params[key];
        this._deflate = zlib.createDeflateRaw(_objectSpread2(_objectSpread2({}, this._options.zlibDeflateOptions), {}, {
          windowBits: windowBits
        }));
        this._deflate[kTotalLength] = 0;
        this._deflate[kBuffers] = [];
        this._deflate.on('data', deflateOnData);
      }
      this._deflate[kCallback] = callback;
      this._deflate.write(data);
      this._deflate.flush(zlib.Z_SYNC_FLUSH, function () {
        if (!_this5._deflate) {
          //
          // The deflate stream was closed while data was being processed.
          //
          return;
        }
        var data = bufferUtil.concat(_this5._deflate[kBuffers], _this5._deflate[kTotalLength]);
        if (fin) {
          data = new FastBuffer$1(data.buffer, data.byteOffset, data.length - 4);
        }

        //
        // Ensure that the callback will not be called again in
        // `PerMessageDeflate#cleanup()`.
        //
        _this5._deflate[kCallback] = null;
        _this5._deflate[kTotalLength] = 0;
        _this5._deflate[kBuffers] = [];
        if (fin && _this5.params["".concat(endpoint, "_no_context_takeover")]) {
          _this5._deflate.reset();
        }
        callback(null, data);
      });
    }
  }], [{
    key: "extensionName",
    get: function get() {
      return 'permessage-deflate';
    }
  }]);
}();
var permessageDeflate = PerMessageDeflate$3;

/**
 * The listener of the `zlib.DeflateRaw` stream `'data'` event.
 *
 * @param {Buffer} chunk A chunk of data
 * @private
 */
function deflateOnData(chunk) {
  this[kBuffers].push(chunk);
  this[kTotalLength] += chunk.length;
}

/**
 * The listener of the `zlib.InflateRaw` stream `'data'` event.
 *
 * @param {Buffer} chunk A chunk of data
 * @private
 */
function inflateOnData(chunk) {
  this[kTotalLength] += chunk.length;
  if (this[kPerMessageDeflate]._maxPayload < 1 || this[kTotalLength] <= this[kPerMessageDeflate]._maxPayload) {
    this[kBuffers].push(chunk);
    return;
  }
  this[kError$1] = new RangeError('Max payload size exceeded');
  this[kError$1].code = 'WS_ERR_UNSUPPORTED_MESSAGE_LENGTH';
  this[kError$1][kStatusCode$2] = 1009;
  this.removeListener('data', inflateOnData);
  this.reset();
}

/**
 * The listener of the `zlib.InflateRaw` stream `'error'` event.
 *
 * @param {Error} err The emitted error
 * @private
 */
function inflateOnError(err) {
  //
  // There is no need to call `Zlib#close()` as the handle is automatically
  // closed when an error is emitted.
  //
  this[kPerMessageDeflate]._inflate = null;
  err[kStatusCode$2] = 1007;
  this[kCallback](err);
}

var validation = {exports: {}};

var isValidUTF8_1;
var isUtf8 = require$$0$3.isUtf8;

//
// Allowed token characters:
//
// '!', '#', '$', '%', '&', ''', '*', '+', '-',
// '.', 0-9, A-Z, '^', '_', '`', a-z, '|', '~'
//
// tokenChars[32] === 0 // ' '
// tokenChars[33] === 1 // '!'
// tokenChars[34] === 0 // '"'
// ...
//
// prettier-ignore
var tokenChars$1 = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 0 - 15
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 16 - 31
0, 1, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 0,
// 32 - 47
1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0,
// 48 - 63
0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
// 64 - 79
1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1,
// 80 - 95
1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
// 96 - 111
1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0 // 112 - 127
];

/**
 * Checks if a status code is allowed in a close frame.
 *
 * @param {Number} code The status code
 * @return {Boolean} `true` if the status code is valid, else `false`
 * @public
 */
function isValidStatusCode$2(code) {
  return code >= 1000 && code <= 1014 && code !== 1004 && code !== 1005 && code !== 1006 || code >= 3000 && code <= 4999;
}

/**
 * Checks if a given buffer contains only correct UTF-8.
 * Ported from https://www.cl.cam.ac.uk/%7Emgk25/ucs/utf8_check.c by
 * Markus Kuhn.
 *
 * @param {Buffer} buf The buffer to check
 * @return {Boolean} `true` if `buf` contains only correct UTF-8, else `false`
 * @public
 */
function _isValidUTF8(buf) {
  var len = buf.length;
  var i = 0;
  while (i < len) {
    if ((buf[i] & 0x80) === 0) {
      // 0xxxxxxx
      i++;
    } else if ((buf[i] & 0xe0) === 0xc0) {
      // 110xxxxx 10xxxxxx
      if (i + 1 === len || (buf[i + 1] & 0xc0) !== 0x80 || (buf[i] & 0xfe) === 0xc0 // Overlong
      ) {
        return false;
      }
      i += 2;
    } else if ((buf[i] & 0xf0) === 0xe0) {
      // 1110xxxx 10xxxxxx 10xxxxxx
      if (i + 2 >= len || (buf[i + 1] & 0xc0) !== 0x80 || (buf[i + 2] & 0xc0) !== 0x80 || buf[i] === 0xe0 && (buf[i + 1] & 0xe0) === 0x80 ||
      // Overlong
      buf[i] === 0xed && (buf[i + 1] & 0xe0) === 0xa0 // Surrogate (U+D800 - U+DFFF)
      ) {
        return false;
      }
      i += 3;
    } else if ((buf[i] & 0xf8) === 0xf0) {
      // 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
      if (i + 3 >= len || (buf[i + 1] & 0xc0) !== 0x80 || (buf[i + 2] & 0xc0) !== 0x80 || (buf[i + 3] & 0xc0) !== 0x80 || buf[i] === 0xf0 && (buf[i + 1] & 0xf0) === 0x80 ||
      // Overlong
      buf[i] === 0xf4 && buf[i + 1] > 0x8f || buf[i] > 0xf4 // > U+10FFFF
      ) {
        return false;
      }
      i += 4;
    } else {
      return false;
    }
  }
  return true;
}
validation.exports = {
  isValidStatusCode: isValidStatusCode$2,
  isValidUTF8: _isValidUTF8,
  tokenChars: tokenChars$1
};
if (isUtf8) {
  isValidUTF8_1 = validation.exports.isValidUTF8 = function (buf) {
    return buf.length < 24 ? _isValidUTF8(buf) : isUtf8(buf);
  };
} /* istanbul ignore else  */else if (!process.env.WS_NO_UTF_8_VALIDATE) {
  try {
    var isValidUTF8$1 = require('utf-8-validate');
    isValidUTF8_1 = validation.exports.isValidUTF8 = function (buf) {
      return buf.length < 32 ? _isValidUTF8(buf) : isValidUTF8$1(buf);
    };
  } catch (e) {
    // Continue regardless of the error.
  }
}
var validationExports = validation.exports;

var Writable = require$$0$1.Writable;
var PerMessageDeflate$2 = permessageDeflate;
var BINARY_TYPES$1 = constants.BINARY_TYPES,
  EMPTY_BUFFER$2 = constants.EMPTY_BUFFER,
  kStatusCode$1 = constants.kStatusCode,
  kWebSocket$1 = constants.kWebSocket;
var concat = bufferUtilExports.concat,
  toArrayBuffer = bufferUtilExports.toArrayBuffer,
  unmask = bufferUtilExports.unmask;
var isValidStatusCode$1 = validationExports.isValidStatusCode,
  isValidUTF8 = validationExports.isValidUTF8;
var FastBuffer = Buffer[Symbol.species];
var GET_INFO = 0;
var GET_PAYLOAD_LENGTH_16 = 1;
var GET_PAYLOAD_LENGTH_64 = 2;
var GET_MASK = 3;
var GET_DATA = 4;
var INFLATING = 5;
var DEFER_EVENT = 6;

/**
 * HyBi Receiver implementation.
 *
 * @extends Writable
 */
var Receiver$1 = /*#__PURE__*/function (_Writable) {
  /**
   * Creates a Receiver instance.
   *
   * @param {Object} [options] Options object
   * @param {Boolean} [options.allowSynchronousEvents=true] Specifies whether
   *     any of the `'message'`, `'ping'`, and `'pong'` events can be emitted
   *     multiple times in the same tick
   * @param {String} [options.binaryType=nodebuffer] The type for binary data
   * @param {Object} [options.extensions] An object containing the negotiated
   *     extensions
   * @param {Boolean} [options.isServer=false] Specifies whether to operate in
   *     client or server mode
   * @param {Number} [options.maxPayload=0] The maximum allowed message length
   * @param {Boolean} [options.skipUTF8Validation=false] Specifies whether or
   *     not to skip UTF-8 validation for text and close messages
   */
  function Receiver() {
    var _this;
    var options = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {};
    _classCallCheck(this, Receiver);
    _this = _callSuper(this, Receiver);
    _this._allowSynchronousEvents = options.allowSynchronousEvents !== undefined ? options.allowSynchronousEvents : true;
    _this._binaryType = options.binaryType || BINARY_TYPES$1[0];
    _this._extensions = options.extensions || {};
    _this._isServer = !!options.isServer;
    _this._maxPayload = options.maxPayload | 0;
    _this._skipUTF8Validation = !!options.skipUTF8Validation;
    _this[kWebSocket$1] = undefined;
    _this._bufferedBytes = 0;
    _this._buffers = [];
    _this._compressed = false;
    _this._payloadLength = 0;
    _this._mask = undefined;
    _this._fragmented = 0;
    _this._masked = false;
    _this._fin = false;
    _this._opcode = 0;
    _this._totalPayloadLength = 0;
    _this._messageLength = 0;
    _this._fragments = [];
    _this._errored = false;
    _this._loop = false;
    _this._state = GET_INFO;
    return _this;
  }

  /**
   * Implements `Writable.prototype._write()`.
   *
   * @param {Buffer} chunk The chunk of data to write
   * @param {String} encoding The character encoding of `chunk`
   * @param {Function} cb Callback
   * @private
   */
  _inherits(Receiver, _Writable);
  return _createClass(Receiver, [{
    key: "_write",
    value: function _write(chunk, encoding, cb) {
      if (this._opcode === 0x08 && this._state == GET_INFO) return cb();
      this._bufferedBytes += chunk.length;
      this._buffers.push(chunk);
      this.startLoop(cb);
    }

    /**
     * Consumes `n` bytes from the buffered data.
     *
     * @param {Number} n The number of bytes to consume
     * @return {Buffer} The consumed bytes
     * @private
     */
  }, {
    key: "consume",
    value: function consume(n) {
      this._bufferedBytes -= n;
      if (n === this._buffers[0].length) return this._buffers.shift();
      if (n < this._buffers[0].length) {
        var buf = this._buffers[0];
        this._buffers[0] = new FastBuffer(buf.buffer, buf.byteOffset + n, buf.length - n);
        return new FastBuffer(buf.buffer, buf.byteOffset, n);
      }
      var dst = Buffer.allocUnsafe(n);
      do {
        var _buf = this._buffers[0];
        var offset = dst.length - n;
        if (n >= _buf.length) {
          dst.set(this._buffers.shift(), offset);
        } else {
          dst.set(new Uint8Array(_buf.buffer, _buf.byteOffset, n), offset);
          this._buffers[0] = new FastBuffer(_buf.buffer, _buf.byteOffset + n, _buf.length - n);
        }
        n -= _buf.length;
      } while (n > 0);
      return dst;
    }

    /**
     * Starts the parsing loop.
     *
     * @param {Function} cb Callback
     * @private
     */
  }, {
    key: "startLoop",
    value: function startLoop(cb) {
      this._loop = true;
      do {
        switch (this._state) {
          case GET_INFO:
            this.getInfo(cb);
            break;
          case GET_PAYLOAD_LENGTH_16:
            this.getPayloadLength16(cb);
            break;
          case GET_PAYLOAD_LENGTH_64:
            this.getPayloadLength64(cb);
            break;
          case GET_MASK:
            this.getMask();
            break;
          case GET_DATA:
            this.getData(cb);
            break;
          case INFLATING:
          case DEFER_EVENT:
            this._loop = false;
            return;
        }
      } while (this._loop);
      if (!this._errored) cb();
    }

    /**
     * Reads the first two bytes of a frame.
     *
     * @param {Function} cb Callback
     * @private
     */
  }, {
    key: "getInfo",
    value: function getInfo(cb) {
      if (this._bufferedBytes < 2) {
        this._loop = false;
        return;
      }
      var buf = this.consume(2);
      if ((buf[0] & 0x30) !== 0x00) {
        var error = this.createError(RangeError, 'RSV2 and RSV3 must be clear', true, 1002, 'WS_ERR_UNEXPECTED_RSV_2_3');
        cb(error);
        return;
      }
      var compressed = (buf[0] & 0x40) === 0x40;
      if (compressed && !this._extensions[PerMessageDeflate$2.extensionName]) {
        var _error = this.createError(RangeError, 'RSV1 must be clear', true, 1002, 'WS_ERR_UNEXPECTED_RSV_1');
        cb(_error);
        return;
      }
      this._fin = (buf[0] & 0x80) === 0x80;
      this._opcode = buf[0] & 0x0f;
      this._payloadLength = buf[1] & 0x7f;
      if (this._opcode === 0x00) {
        if (compressed) {
          var _error2 = this.createError(RangeError, 'RSV1 must be clear', true, 1002, 'WS_ERR_UNEXPECTED_RSV_1');
          cb(_error2);
          return;
        }
        if (!this._fragmented) {
          var _error3 = this.createError(RangeError, 'invalid opcode 0', true, 1002, 'WS_ERR_INVALID_OPCODE');
          cb(_error3);
          return;
        }
        this._opcode = this._fragmented;
      } else if (this._opcode === 0x01 || this._opcode === 0x02) {
        if (this._fragmented) {
          var _error4 = this.createError(RangeError, "invalid opcode ".concat(this._opcode), true, 1002, 'WS_ERR_INVALID_OPCODE');
          cb(_error4);
          return;
        }
        this._compressed = compressed;
      } else if (this._opcode > 0x07 && this._opcode < 0x0b) {
        if (!this._fin) {
          var _error5 = this.createError(RangeError, 'FIN must be set', true, 1002, 'WS_ERR_EXPECTED_FIN');
          cb(_error5);
          return;
        }
        if (compressed) {
          var _error6 = this.createError(RangeError, 'RSV1 must be clear', true, 1002, 'WS_ERR_UNEXPECTED_RSV_1');
          cb(_error6);
          return;
        }
        if (this._payloadLength > 0x7d || this._opcode === 0x08 && this._payloadLength === 1) {
          var _error7 = this.createError(RangeError, "invalid payload length ".concat(this._payloadLength), true, 1002, 'WS_ERR_INVALID_CONTROL_PAYLOAD_LENGTH');
          cb(_error7);
          return;
        }
      } else {
        var _error8 = this.createError(RangeError, "invalid opcode ".concat(this._opcode), true, 1002, 'WS_ERR_INVALID_OPCODE');
        cb(_error8);
        return;
      }
      if (!this._fin && !this._fragmented) this._fragmented = this._opcode;
      this._masked = (buf[1] & 0x80) === 0x80;
      if (this._isServer) {
        if (!this._masked) {
          var _error9 = this.createError(RangeError, 'MASK must be set', true, 1002, 'WS_ERR_EXPECTED_MASK');
          cb(_error9);
          return;
        }
      } else if (this._masked) {
        var _error10 = this.createError(RangeError, 'MASK must be clear', true, 1002, 'WS_ERR_UNEXPECTED_MASK');
        cb(_error10);
        return;
      }
      if (this._payloadLength === 126) this._state = GET_PAYLOAD_LENGTH_16;else if (this._payloadLength === 127) this._state = GET_PAYLOAD_LENGTH_64;else this.haveLength(cb);
    }

    /**
     * Gets extended payload length (7+16).
     *
     * @param {Function} cb Callback
     * @private
     */
  }, {
    key: "getPayloadLength16",
    value: function getPayloadLength16(cb) {
      if (this._bufferedBytes < 2) {
        this._loop = false;
        return;
      }
      this._payloadLength = this.consume(2).readUInt16BE(0);
      this.haveLength(cb);
    }

    /**
     * Gets extended payload length (7+64).
     *
     * @param {Function} cb Callback
     * @private
     */
  }, {
    key: "getPayloadLength64",
    value: function getPayloadLength64(cb) {
      if (this._bufferedBytes < 8) {
        this._loop = false;
        return;
      }
      var buf = this.consume(8);
      var num = buf.readUInt32BE(0);

      //
      // The maximum safe integer in JavaScript is 2^53 - 1. An error is returned
      // if payload length is greater than this number.
      //
      if (num > Math.pow(2, 53 - 32) - 1) {
        var error = this.createError(RangeError, 'Unsupported WebSocket frame: payload length > 2^53 - 1', false, 1009, 'WS_ERR_UNSUPPORTED_DATA_PAYLOAD_LENGTH');
        cb(error);
        return;
      }
      this._payloadLength = num * Math.pow(2, 32) + buf.readUInt32BE(4);
      this.haveLength(cb);
    }

    /**
     * Payload length has been read.
     *
     * @param {Function} cb Callback
     * @private
     */
  }, {
    key: "haveLength",
    value: function haveLength(cb) {
      if (this._payloadLength && this._opcode < 0x08) {
        this._totalPayloadLength += this._payloadLength;
        if (this._totalPayloadLength > this._maxPayload && this._maxPayload > 0) {
          var error = this.createError(RangeError, 'Max payload size exceeded', false, 1009, 'WS_ERR_UNSUPPORTED_MESSAGE_LENGTH');
          cb(error);
          return;
        }
      }
      if (this._masked) this._state = GET_MASK;else this._state = GET_DATA;
    }

    /**
     * Reads mask bytes.
     *
     * @private
     */
  }, {
    key: "getMask",
    value: function getMask() {
      if (this._bufferedBytes < 4) {
        this._loop = false;
        return;
      }
      this._mask = this.consume(4);
      this._state = GET_DATA;
    }

    /**
     * Reads data bytes.
     *
     * @param {Function} cb Callback
     * @private
     */
  }, {
    key: "getData",
    value: function getData(cb) {
      var data = EMPTY_BUFFER$2;
      if (this._payloadLength) {
        if (this._bufferedBytes < this._payloadLength) {
          this._loop = false;
          return;
        }
        data = this.consume(this._payloadLength);
        if (this._masked && (this._mask[0] | this._mask[1] | this._mask[2] | this._mask[3]) !== 0) {
          unmask(data, this._mask);
        }
      }
      if (this._opcode > 0x07) {
        this.controlMessage(data, cb);
        return;
      }
      if (this._compressed) {
        this._state = INFLATING;
        this.decompress(data, cb);
        return;
      }
      if (data.length) {
        //
        // This message is not compressed so its length is the sum of the payload
        // length of all fragments.
        //
        this._messageLength = this._totalPayloadLength;
        this._fragments.push(data);
      }
      this.dataMessage(cb);
    }

    /**
     * Decompresses data.
     *
     * @param {Buffer} data Compressed data
     * @param {Function} cb Callback
     * @private
     */
  }, {
    key: "decompress",
    value: function decompress(data, cb) {
      var _this2 = this;
      var perMessageDeflate = this._extensions[PerMessageDeflate$2.extensionName];
      perMessageDeflate.decompress(data, this._fin, function (err, buf) {
        if (err) return cb(err);
        if (buf.length) {
          _this2._messageLength += buf.length;
          if (_this2._messageLength > _this2._maxPayload && _this2._maxPayload > 0) {
            var error = _this2.createError(RangeError, 'Max payload size exceeded', false, 1009, 'WS_ERR_UNSUPPORTED_MESSAGE_LENGTH');
            cb(error);
            return;
          }
          _this2._fragments.push(buf);
        }
        _this2.dataMessage(cb);
        if (_this2._state === GET_INFO) _this2.startLoop(cb);
      });
    }

    /**
     * Handles a data message.
     *
     * @param {Function} cb Callback
     * @private
     */
  }, {
    key: "dataMessage",
    value: function dataMessage(cb) {
      var _this3 = this;
      if (!this._fin) {
        this._state = GET_INFO;
        return;
      }
      var messageLength = this._messageLength;
      var fragments = this._fragments;
      this._totalPayloadLength = 0;
      this._messageLength = 0;
      this._fragmented = 0;
      this._fragments = [];
      if (this._opcode === 2) {
        var data;
        if (this._binaryType === 'nodebuffer') {
          data = concat(fragments, messageLength);
        } else if (this._binaryType === 'arraybuffer') {
          data = toArrayBuffer(concat(fragments, messageLength));
        } else {
          data = fragments;
        }
        if (this._allowSynchronousEvents) {
          this.emit('message', data, true);
          this._state = GET_INFO;
        } else {
          this._state = DEFER_EVENT;
          setImmediate(function () {
            _this3.emit('message', data, true);
            _this3._state = GET_INFO;
            _this3.startLoop(cb);
          });
        }
      } else {
        var buf = concat(fragments, messageLength);
        if (!this._skipUTF8Validation && !isValidUTF8(buf)) {
          var error = this.createError(Error, 'invalid UTF-8 sequence', true, 1007, 'WS_ERR_INVALID_UTF8');
          cb(error);
          return;
        }
        if (this._state === INFLATING || this._allowSynchronousEvents) {
          this.emit('message', buf, false);
          this._state = GET_INFO;
        } else {
          this._state = DEFER_EVENT;
          setImmediate(function () {
            _this3.emit('message', buf, false);
            _this3._state = GET_INFO;
            _this3.startLoop(cb);
          });
        }
      }
    }

    /**
     * Handles a control message.
     *
     * @param {Buffer} data Data to handle
     * @return {(Error|RangeError|undefined)} A possible error
     * @private
     */
  }, {
    key: "controlMessage",
    value: function controlMessage(data, cb) {
      var _this4 = this;
      if (this._opcode === 0x08) {
        if (data.length === 0) {
          this._loop = false;
          this.emit('conclude', 1005, EMPTY_BUFFER$2);
          this.end();
        } else {
          var code = data.readUInt16BE(0);
          if (!isValidStatusCode$1(code)) {
            var error = this.createError(RangeError, "invalid status code ".concat(code), true, 1002, 'WS_ERR_INVALID_CLOSE_CODE');
            cb(error);
            return;
          }
          var buf = new FastBuffer(data.buffer, data.byteOffset + 2, data.length - 2);
          if (!this._skipUTF8Validation && !isValidUTF8(buf)) {
            var _error11 = this.createError(Error, 'invalid UTF-8 sequence', true, 1007, 'WS_ERR_INVALID_UTF8');
            cb(_error11);
            return;
          }
          this._loop = false;
          this.emit('conclude', code, buf);
          this.end();
        }
        this._state = GET_INFO;
        return;
      }
      if (this._allowSynchronousEvents) {
        this.emit(this._opcode === 0x09 ? 'ping' : 'pong', data);
        this._state = GET_INFO;
      } else {
        this._state = DEFER_EVENT;
        setImmediate(function () {
          _this4.emit(_this4._opcode === 0x09 ? 'ping' : 'pong', data);
          _this4._state = GET_INFO;
          _this4.startLoop(cb);
        });
      }
    }

    /**
     * Builds an error object.
     *
     * @param {function(new:Error|RangeError)} ErrorCtor The error constructor
     * @param {String} message The error message
     * @param {Boolean} prefix Specifies whether or not to add a default prefix to
     *     `message`
     * @param {Number} statusCode The status code
     * @param {String} errorCode The exposed error code
     * @return {(Error|RangeError)} The error
     * @private
     */
  }, {
    key: "createError",
    value: function createError(ErrorCtor, message, prefix, statusCode, errorCode) {
      this._loop = false;
      this._errored = true;
      var err = new ErrorCtor(prefix ? "Invalid WebSocket frame: ".concat(message) : message);
      Error.captureStackTrace(err, this.createError);
      err.code = errorCode;
      err[kStatusCode$1] = statusCode;
      return err;
    }
  }]);
}(Writable);
var receiver = Receiver$1;

require$$0$1.Duplex;
var randomFillSync = require$$5.randomFillSync;
var PerMessageDeflate$1 = permessageDeflate;
var EMPTY_BUFFER$1 = constants.EMPTY_BUFFER;
var isValidStatusCode = validationExports.isValidStatusCode;
var applyMask = bufferUtilExports.mask,
  toBuffer$1 = bufferUtilExports.toBuffer;
var kByteLength = Symbol('kByteLength');
var maskBuffer = Buffer.alloc(4);
var RANDOM_POOL_SIZE = 8 * 1024;
var randomPool;
var randomPoolPointer = RANDOM_POOL_SIZE;

/**
 * HyBi Sender implementation.
 */
var Sender$1 = /*#__PURE__*/function () {
  /**
   * Creates a Sender instance.
   *
   * @param {Duplex} socket The connection socket
   * @param {Object} [extensions] An object containing the negotiated extensions
   * @param {Function} [generateMask] The function used to generate the masking
   *     key
   */
  function Sender(socket, extensions, generateMask) {
    _classCallCheck(this, Sender);
    this._extensions = extensions || {};
    if (generateMask) {
      this._generateMask = generateMask;
      this._maskBuffer = Buffer.alloc(4);
    }
    this._socket = socket;
    this._firstFragment = true;
    this._compress = false;
    this._bufferedBytes = 0;
    this._deflating = false;
    this._queue = [];
  }

  /**
   * Frames a piece of data according to the HyBi WebSocket protocol.
   *
   * @param {(Buffer|String)} data The data to frame
   * @param {Object} options Options object
   * @param {Boolean} [options.fin=false] Specifies whether or not to set the
   *     FIN bit
   * @param {Function} [options.generateMask] The function used to generate the
   *     masking key
   * @param {Boolean} [options.mask=false] Specifies whether or not to mask
   *     `data`
   * @param {Buffer} [options.maskBuffer] The buffer used to store the masking
   *     key
   * @param {Number} options.opcode The opcode
   * @param {Boolean} [options.readOnly=false] Specifies whether `data` can be
   *     modified
   * @param {Boolean} [options.rsv1=false] Specifies whether or not to set the
   *     RSV1 bit
   * @return {(Buffer|String)[]} The framed data
   * @public
   */
  return _createClass(Sender, [{
    key: "close",
    value:
    /**
     * Sends a close message to the other peer.
     *
     * @param {Number} [code] The status code component of the body
     * @param {(String|Buffer)} [data] The message component of the body
     * @param {Boolean} [mask=false] Specifies whether or not to mask the message
     * @param {Function} [cb] Callback
     * @public
     */
    function close(code, data, mask, cb) {
      var buf;
      if (code === undefined) {
        buf = EMPTY_BUFFER$1;
      } else if (typeof code !== 'number' || !isValidStatusCode(code)) {
        throw new TypeError('First argument must be a valid error code number');
      } else if (data === undefined || !data.length) {
        buf = Buffer.allocUnsafe(2);
        buf.writeUInt16BE(code, 0);
      } else {
        var length = Buffer.byteLength(data);
        if (length > 123) {
          throw new RangeError('The message must not be greater than 123 bytes');
        }
        buf = Buffer.allocUnsafe(2 + length);
        buf.writeUInt16BE(code, 0);
        if (typeof data === 'string') {
          buf.write(data, 2);
        } else {
          buf.set(data, 2);
        }
      }
      var options = _defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty({}, kByteLength, buf.length), "fin", true), "generateMask", this._generateMask), "mask", mask), "maskBuffer", this._maskBuffer), "opcode", 0x08), "readOnly", false), "rsv1", false);
      if (this._deflating) {
        this.enqueue([this.dispatch, buf, false, options, cb]);
      } else {
        this.sendFrame(Sender.frame(buf, options), cb);
      }
    }

    /**
     * Sends a ping message to the other peer.
     *
     * @param {*} data The message to send
     * @param {Boolean} [mask=false] Specifies whether or not to mask `data`
     * @param {Function} [cb] Callback
     * @public
     */
  }, {
    key: "ping",
    value: function ping(data, mask, cb) {
      var byteLength;
      var readOnly;
      if (typeof data === 'string') {
        byteLength = Buffer.byteLength(data);
        readOnly = false;
      } else {
        data = toBuffer$1(data);
        byteLength = data.length;
        readOnly = toBuffer$1.readOnly;
      }
      if (byteLength > 125) {
        throw new RangeError('The data size must not be greater than 125 bytes');
      }
      var options = _defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty({}, kByteLength, byteLength), "fin", true), "generateMask", this._generateMask), "mask", mask), "maskBuffer", this._maskBuffer), "opcode", 0x09), "readOnly", readOnly), "rsv1", false);
      if (this._deflating) {
        this.enqueue([this.dispatch, data, false, options, cb]);
      } else {
        this.sendFrame(Sender.frame(data, options), cb);
      }
    }

    /**
     * Sends a pong message to the other peer.
     *
     * @param {*} data The message to send
     * @param {Boolean} [mask=false] Specifies whether or not to mask `data`
     * @param {Function} [cb] Callback
     * @public
     */
  }, {
    key: "pong",
    value: function pong(data, mask, cb) {
      var byteLength;
      var readOnly;
      if (typeof data === 'string') {
        byteLength = Buffer.byteLength(data);
        readOnly = false;
      } else {
        data = toBuffer$1(data);
        byteLength = data.length;
        readOnly = toBuffer$1.readOnly;
      }
      if (byteLength > 125) {
        throw new RangeError('The data size must not be greater than 125 bytes');
      }
      var options = _defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty({}, kByteLength, byteLength), "fin", true), "generateMask", this._generateMask), "mask", mask), "maskBuffer", this._maskBuffer), "opcode", 0x0a), "readOnly", readOnly), "rsv1", false);
      if (this._deflating) {
        this.enqueue([this.dispatch, data, false, options, cb]);
      } else {
        this.sendFrame(Sender.frame(data, options), cb);
      }
    }

    /**
     * Sends a data message to the other peer.
     *
     * @param {*} data The message to send
     * @param {Object} options Options object
     * @param {Boolean} [options.binary=false] Specifies whether `data` is binary
     *     or text
     * @param {Boolean} [options.compress=false] Specifies whether or not to
     *     compress `data`
     * @param {Boolean} [options.fin=false] Specifies whether the fragment is the
     *     last one
     * @param {Boolean} [options.mask=false] Specifies whether or not to mask
     *     `data`
     * @param {Function} [cb] Callback
     * @public
     */
  }, {
    key: "send",
    value: function send(data, options, cb) {
      var perMessageDeflate = this._extensions[PerMessageDeflate$1.extensionName];
      var opcode = options.binary ? 2 : 1;
      var rsv1 = options.compress;
      var byteLength;
      var readOnly;
      if (typeof data === 'string') {
        byteLength = Buffer.byteLength(data);
        readOnly = false;
      } else {
        data = toBuffer$1(data);
        byteLength = data.length;
        readOnly = toBuffer$1.readOnly;
      }
      if (this._firstFragment) {
        this._firstFragment = false;
        if (rsv1 && perMessageDeflate && perMessageDeflate.params[perMessageDeflate._isServer ? 'server_no_context_takeover' : 'client_no_context_takeover']) {
          rsv1 = byteLength >= perMessageDeflate._threshold;
        }
        this._compress = rsv1;
      } else {
        rsv1 = false;
        opcode = 0;
      }
      if (options.fin) this._firstFragment = true;
      if (perMessageDeflate) {
        var opts = _defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty({}, kByteLength, byteLength), "fin", options.fin), "generateMask", this._generateMask), "mask", options.mask), "maskBuffer", this._maskBuffer), "opcode", opcode), "readOnly", readOnly), "rsv1", rsv1);
        if (this._deflating) {
          this.enqueue([this.dispatch, data, this._compress, opts, cb]);
        } else {
          this.dispatch(data, this._compress, opts, cb);
        }
      } else {
        this.sendFrame(Sender.frame(data, _defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty(_defineProperty({}, kByteLength, byteLength), "fin", options.fin), "generateMask", this._generateMask), "mask", options.mask), "maskBuffer", this._maskBuffer), "opcode", opcode), "readOnly", readOnly), "rsv1", false)), cb);
      }
    }

    /**
     * Dispatches a message.
     *
     * @param {(Buffer|String)} data The message to send
     * @param {Boolean} [compress=false] Specifies whether or not to compress
     *     `data`
     * @param {Object} options Options object
     * @param {Boolean} [options.fin=false] Specifies whether or not to set the
     *     FIN bit
     * @param {Function} [options.generateMask] The function used to generate the
     *     masking key
     * @param {Boolean} [options.mask=false] Specifies whether or not to mask
     *     `data`
     * @param {Buffer} [options.maskBuffer] The buffer used to store the masking
     *     key
     * @param {Number} options.opcode The opcode
     * @param {Boolean} [options.readOnly=false] Specifies whether `data` can be
     *     modified
     * @param {Boolean} [options.rsv1=false] Specifies whether or not to set the
     *     RSV1 bit
     * @param {Function} [cb] Callback
     * @private
     */
  }, {
    key: "dispatch",
    value: function dispatch(data, compress, options, cb) {
      var _this = this;
      if (!compress) {
        this.sendFrame(Sender.frame(data, options), cb);
        return;
      }
      var perMessageDeflate = this._extensions[PerMessageDeflate$1.extensionName];
      this._bufferedBytes += options[kByteLength];
      this._deflating = true;
      perMessageDeflate.compress(data, options.fin, function (_, buf) {
        if (_this._socket.destroyed) {
          var err = new Error('The socket was closed while data was being compressed');
          if (typeof cb === 'function') cb(err);
          for (var i = 0; i < _this._queue.length; i++) {
            var params = _this._queue[i];
            var callback = params[params.length - 1];
            if (typeof callback === 'function') callback(err);
          }
          return;
        }
        _this._bufferedBytes -= options[kByteLength];
        _this._deflating = false;
        options.readOnly = false;
        _this.sendFrame(Sender.frame(buf, options), cb);
        _this.dequeue();
      });
    }

    /**
     * Executes queued send operations.
     *
     * @private
     */
  }, {
    key: "dequeue",
    value: function dequeue() {
      while (!this._deflating && this._queue.length) {
        var params = this._queue.shift();
        this._bufferedBytes -= params[3][kByteLength];
        Reflect.apply(params[0], this, params.slice(1));
      }
    }

    /**
     * Enqueues a send operation.
     *
     * @param {Array} params Send operation parameters.
     * @private
     */
  }, {
    key: "enqueue",
    value: function enqueue(params) {
      this._bufferedBytes += params[3][kByteLength];
      this._queue.push(params);
    }

    /**
     * Sends a frame.
     *
     * @param {Buffer[]} list The frame to send
     * @param {Function} [cb] Callback
     * @private
     */
  }, {
    key: "sendFrame",
    value: function sendFrame(list, cb) {
      if (list.length === 2) {
        this._socket.cork();
        this._socket.write(list[0]);
        this._socket.write(list[1], cb);
        this._socket.uncork();
      } else {
        this._socket.write(list[0], cb);
      }
    }
  }], [{
    key: "frame",
    value: function frame(data, options) {
      var mask;
      var merge = false;
      var offset = 2;
      var skipMasking = false;
      if (options.mask) {
        mask = options.maskBuffer || maskBuffer;
        if (options.generateMask) {
          options.generateMask(mask);
        } else {
          if (randomPoolPointer === RANDOM_POOL_SIZE) {
            /* istanbul ignore else  */
            if (randomPool === undefined) {
              //
              // This is lazily initialized because server-sent frames must not
              // be masked so it may never be used.
              //
              randomPool = Buffer.alloc(RANDOM_POOL_SIZE);
            }
            randomFillSync(randomPool, 0, RANDOM_POOL_SIZE);
            randomPoolPointer = 0;
          }
          mask[0] = randomPool[randomPoolPointer++];
          mask[1] = randomPool[randomPoolPointer++];
          mask[2] = randomPool[randomPoolPointer++];
          mask[3] = randomPool[randomPoolPointer++];
        }
        skipMasking = (mask[0] | mask[1] | mask[2] | mask[3]) === 0;
        offset = 6;
      }
      var dataLength;
      if (typeof data === 'string') {
        if ((!options.mask || skipMasking) && options[kByteLength] !== undefined) {
          dataLength = options[kByteLength];
        } else {
          data = Buffer.from(data);
          dataLength = data.length;
        }
      } else {
        dataLength = data.length;
        merge = options.mask && options.readOnly && !skipMasking;
      }
      var payloadLength = dataLength;
      if (dataLength >= 65536) {
        offset += 8;
        payloadLength = 127;
      } else if (dataLength > 125) {
        offset += 2;
        payloadLength = 126;
      }
      var target = Buffer.allocUnsafe(merge ? dataLength + offset : offset);
      target[0] = options.fin ? options.opcode | 0x80 : options.opcode;
      if (options.rsv1) target[0] |= 0x40;
      target[1] = payloadLength;
      if (payloadLength === 126) {
        target.writeUInt16BE(dataLength, 2);
      } else if (payloadLength === 127) {
        target[2] = target[3] = 0;
        target.writeUIntBE(dataLength, 4, 6);
      }
      if (!options.mask) return [target, data];
      target[1] |= 0x80;
      target[offset - 4] = mask[0];
      target[offset - 3] = mask[1];
      target[offset - 2] = mask[2];
      target[offset - 1] = mask[3];
      if (skipMasking) return [target, data];
      if (merge) {
        applyMask(data, mask, target, offset, dataLength);
        return [target];
      }
      applyMask(data, mask, data, 0, dataLength);
      return [target, data];
    }
  }]);
}();
var sender = Sender$1;

var kForOnEventAttribute$1 = constants.kForOnEventAttribute,
  kListener$1 = constants.kListener;
var kCode = Symbol('kCode');
var kData = Symbol('kData');
var kError = Symbol('kError');
var kMessage = Symbol('kMessage');
var kReason = Symbol('kReason');
var kTarget = Symbol('kTarget');
var kType = Symbol('kType');
var kWasClean = Symbol('kWasClean');

/**
 * Class representing an event.
 */
var Event = /*#__PURE__*/function () {
  /**
   * Create a new `Event`.
   *
   * @param {String} type The name of the event
   * @throws {TypeError} If the `type` argument is not specified
   */
  function Event(type) {
    _classCallCheck(this, Event);
    this[kTarget] = null;
    this[kType] = type;
  }

  /**
   * @type {*}
   */
  return _createClass(Event, [{
    key: "target",
    get: function get() {
      return this[kTarget];
    }

    /**
     * @type {String}
     */
  }, {
    key: "type",
    get: function get() {
      return this[kType];
    }
  }]);
}();
Object.defineProperty(Event.prototype, 'target', {
  enumerable: true
});
Object.defineProperty(Event.prototype, 'type', {
  enumerable: true
});

/**
 * Class representing a close event.
 *
 * @extends Event
 */
var CloseEvent = /*#__PURE__*/function (_Event) {
  /**
   * Create a new `CloseEvent`.
   *
   * @param {String} type The name of the event
   * @param {Object} [options] A dictionary object that allows for setting
   *     attributes via object members of the same name
   * @param {Number} [options.code=0] The status code explaining why the
   *     connection was closed
   * @param {String} [options.reason=''] A human-readable string explaining why
   *     the connection was closed
   * @param {Boolean} [options.wasClean=false] Indicates whether or not the
   *     connection was cleanly closed
   */
  function CloseEvent(type) {
    var _this;
    var options = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : {};
    _classCallCheck(this, CloseEvent);
    _this = _callSuper(this, CloseEvent, [type]);
    _this[kCode] = options.code === undefined ? 0 : options.code;
    _this[kReason] = options.reason === undefined ? '' : options.reason;
    _this[kWasClean] = options.wasClean === undefined ? false : options.wasClean;
    return _this;
  }

  /**
   * @type {Number}
   */
  _inherits(CloseEvent, _Event);
  return _createClass(CloseEvent, [{
    key: "code",
    get: function get() {
      return this[kCode];
    }

    /**
     * @type {String}
     */
  }, {
    key: "reason",
    get: function get() {
      return this[kReason];
    }

    /**
     * @type {Boolean}
     */
  }, {
    key: "wasClean",
    get: function get() {
      return this[kWasClean];
    }
  }]);
}(Event);
Object.defineProperty(CloseEvent.prototype, 'code', {
  enumerable: true
});
Object.defineProperty(CloseEvent.prototype, 'reason', {
  enumerable: true
});
Object.defineProperty(CloseEvent.prototype, 'wasClean', {
  enumerable: true
});

/**
 * Class representing an error event.
 *
 * @extends Event
 */
var ErrorEvent = /*#__PURE__*/function (_Event2) {
  /**
   * Create a new `ErrorEvent`.
   *
   * @param {String} type The name of the event
   * @param {Object} [options] A dictionary object that allows for setting
   *     attributes via object members of the same name
   * @param {*} [options.error=null] The error that generated this event
   * @param {String} [options.message=''] The error message
   */
  function ErrorEvent(type) {
    var _this2;
    var options = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : {};
    _classCallCheck(this, ErrorEvent);
    _this2 = _callSuper(this, ErrorEvent, [type]);
    _this2[kError] = options.error === undefined ? null : options.error;
    _this2[kMessage] = options.message === undefined ? '' : options.message;
    return _this2;
  }

  /**
   * @type {*}
   */
  _inherits(ErrorEvent, _Event2);
  return _createClass(ErrorEvent, [{
    key: "error",
    get: function get() {
      return this[kError];
    }

    /**
     * @type {String}
     */
  }, {
    key: "message",
    get: function get() {
      return this[kMessage];
    }
  }]);
}(Event);
Object.defineProperty(ErrorEvent.prototype, 'error', {
  enumerable: true
});
Object.defineProperty(ErrorEvent.prototype, 'message', {
  enumerable: true
});

/**
 * Class representing a message event.
 *
 * @extends Event
 */
var MessageEvent = /*#__PURE__*/function (_Event3) {
  /**
   * Create a new `MessageEvent`.
   *
   * @param {String} type The name of the event
   * @param {Object} [options] A dictionary object that allows for setting
   *     attributes via object members of the same name
   * @param {*} [options.data=null] The message content
   */
  function MessageEvent(type) {
    var _this3;
    var options = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : {};
    _classCallCheck(this, MessageEvent);
    _this3 = _callSuper(this, MessageEvent, [type]);
    _this3[kData] = options.data === undefined ? null : options.data;
    return _this3;
  }

  /**
   * @type {*}
   */
  _inherits(MessageEvent, _Event3);
  return _createClass(MessageEvent, [{
    key: "data",
    get: function get() {
      return this[kData];
    }
  }]);
}(Event);
Object.defineProperty(MessageEvent.prototype, 'data', {
  enumerable: true
});

/**
 * This provides methods for emulating the `EventTarget` interface. It's not
 * meant to be used directly.
 *
 * @mixin
 */
var EventTarget = {
  /**
   * Register an event listener.
   *
   * @param {String} type A string representing the event type to listen for
   * @param {(Function|Object)} handler The listener to add
   * @param {Object} [options] An options object specifies characteristics about
   *     the event listener
   * @param {Boolean} [options.once=false] A `Boolean` indicating that the
   *     listener should be invoked at most once after being added. If `true`,
   *     the listener would be automatically removed when invoked.
   * @public
   */
  addEventListener: function addEventListener(type, handler) {
    var options = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : {};
    var _iterator = _createForOfIteratorHelper(this.listeners(type)),
      _step;
    try {
      for (_iterator.s(); !(_step = _iterator.n()).done;) {
        var listener = _step.value;
        if (!options[kForOnEventAttribute$1] && listener[kListener$1] === handler && !listener[kForOnEventAttribute$1]) {
          return;
        }
      }
    } catch (err) {
      _iterator.e(err);
    } finally {
      _iterator.f();
    }
    var wrapper;
    if (type === 'message') {
      wrapper = function onMessage(data, isBinary) {
        var event = new MessageEvent('message', {
          data: isBinary ? data : data.toString()
        });
        event[kTarget] = this;
        callListener(handler, this, event);
      };
    } else if (type === 'close') {
      wrapper = function onClose(code, message) {
        var event = new CloseEvent('close', {
          code: code,
          reason: message.toString(),
          wasClean: this._closeFrameReceived && this._closeFrameSent
        });
        event[kTarget] = this;
        callListener(handler, this, event);
      };
    } else if (type === 'error') {
      wrapper = function onError(error) {
        var event = new ErrorEvent('error', {
          error: error,
          message: error.message
        });
        event[kTarget] = this;
        callListener(handler, this, event);
      };
    } else if (type === 'open') {
      wrapper = function onOpen() {
        var event = new Event('open');
        event[kTarget] = this;
        callListener(handler, this, event);
      };
    } else {
      return;
    }
    wrapper[kForOnEventAttribute$1] = !!options[kForOnEventAttribute$1];
    wrapper[kListener$1] = handler;
    if (options.once) {
      this.once(type, wrapper);
    } else {
      this.on(type, wrapper);
    }
  },
  /**
   * Remove an event listener.
   *
   * @param {String} type A string representing the event type to remove
   * @param {(Function|Object)} handler The listener to remove
   * @public
   */
  removeEventListener: function removeEventListener(type, handler) {
    var _iterator2 = _createForOfIteratorHelper(this.listeners(type)),
      _step2;
    try {
      for (_iterator2.s(); !(_step2 = _iterator2.n()).done;) {
        var listener = _step2.value;
        if (listener[kListener$1] === handler && !listener[kForOnEventAttribute$1]) {
          this.removeListener(type, listener);
          break;
        }
      }
    } catch (err) {
      _iterator2.e(err);
    } finally {
      _iterator2.f();
    }
  }
};
var eventTarget = {
  CloseEvent: CloseEvent,
  ErrorEvent: ErrorEvent,
  Event: Event,
  EventTarget: EventTarget,
  MessageEvent: MessageEvent
};

/**
 * Call an event listener
 *
 * @param {(Function|Object)} listener The listener to call
 * @param {*} thisArg The value to use as `this`` when calling the listener
 * @param {Event} event The event to pass to the listener
 * @private
 */
function callListener(listener, thisArg, event) {
  if (_typeof(listener) === 'object' && listener.handleEvent) {
    listener.handleEvent.call(listener, event);
  } else {
    listener.call(thisArg, event);
  }
}

var tokenChars = validationExports.tokenChars;

/**
 * Adds an offer to the map of extension offers or a parameter to the map of
 * parameters.
 *
 * @param {Object} dest The map of extension offers or parameters
 * @param {String} name The extension or parameter name
 * @param {(Object|Boolean|String)} elem The extension parameters or the
 *     parameter value
 * @private
 */
function push(dest, name, elem) {
  if (dest[name] === undefined) dest[name] = [elem];else dest[name].push(elem);
}

/**
 * Parses the `Sec-WebSocket-Extensions` header into an object.
 *
 * @param {String} header The field value of the header
 * @return {Object} The parsed object
 * @public
 */
function parse$2(header) {
  var offers = Object.create(null);
  var params = Object.create(null);
  var mustUnescape = false;
  var isEscaping = false;
  var inQuotes = false;
  var extensionName;
  var paramName;
  var start = -1;
  var code = -1;
  var end = -1;
  var i = 0;
  for (; i < header.length; i++) {
    code = header.charCodeAt(i);
    if (extensionName === undefined) {
      if (end === -1 && tokenChars[code] === 1) {
        if (start === -1) start = i;
      } else if (i !== 0 && (code === 0x20 /* ' ' */ || code === 0x09) /* '\t' */) {
        if (end === -1 && start !== -1) end = i;
      } else if (code === 0x3b /* ';' */ || code === 0x2c /* ',' */) {
        if (start === -1) {
          throw new SyntaxError("Unexpected character at index ".concat(i));
        }
        if (end === -1) end = i;
        var name = header.slice(start, end);
        if (code === 0x2c) {
          push(offers, name, params);
          params = Object.create(null);
        } else {
          extensionName = name;
        }
        start = end = -1;
      } else {
        throw new SyntaxError("Unexpected character at index ".concat(i));
      }
    } else if (paramName === undefined) {
      if (end === -1 && tokenChars[code] === 1) {
        if (start === -1) start = i;
      } else if (code === 0x20 || code === 0x09) {
        if (end === -1 && start !== -1) end = i;
      } else if (code === 0x3b || code === 0x2c) {
        if (start === -1) {
          throw new SyntaxError("Unexpected character at index ".concat(i));
        }
        if (end === -1) end = i;
        push(params, header.slice(start, end), true);
        if (code === 0x2c) {
          push(offers, extensionName, params);
          params = Object.create(null);
          extensionName = undefined;
        }
        start = end = -1;
      } else if (code === 0x3d /* '=' */ && start !== -1 && end === -1) {
        paramName = header.slice(start, i);
        start = end = -1;
      } else {
        throw new SyntaxError("Unexpected character at index ".concat(i));
      }
    } else {
      //
      // The value of a quoted-string after unescaping must conform to the
      // token ABNF, so only token characters are valid.
      // Ref: https://tools.ietf.org/html/rfc6455#section-9.1
      //
      if (isEscaping) {
        if (tokenChars[code] !== 1) {
          throw new SyntaxError("Unexpected character at index ".concat(i));
        }
        if (start === -1) start = i;else if (!mustUnescape) mustUnescape = true;
        isEscaping = false;
      } else if (inQuotes) {
        if (tokenChars[code] === 1) {
          if (start === -1) start = i;
        } else if (code === 0x22 /* '"' */ && start !== -1) {
          inQuotes = false;
          end = i;
        } else if (code === 0x5c /* '\' */) {
          isEscaping = true;
        } else {
          throw new SyntaxError("Unexpected character at index ".concat(i));
        }
      } else if (code === 0x22 && header.charCodeAt(i - 1) === 0x3d) {
        inQuotes = true;
      } else if (end === -1 && tokenChars[code] === 1) {
        if (start === -1) start = i;
      } else if (start !== -1 && (code === 0x20 || code === 0x09)) {
        if (end === -1) end = i;
      } else if (code === 0x3b || code === 0x2c) {
        if (start === -1) {
          throw new SyntaxError("Unexpected character at index ".concat(i));
        }
        if (end === -1) end = i;
        var value = header.slice(start, end);
        if (mustUnescape) {
          value = value.replace(/\\/g, '');
          mustUnescape = false;
        }
        push(params, paramName, value);
        if (code === 0x2c) {
          push(offers, extensionName, params);
          params = Object.create(null);
          extensionName = undefined;
        }
        paramName = undefined;
        start = end = -1;
      } else {
        throw new SyntaxError("Unexpected character at index ".concat(i));
      }
    }
  }
  if (start === -1 || inQuotes || code === 0x20 || code === 0x09) {
    throw new SyntaxError('Unexpected end of input');
  }
  if (end === -1) end = i;
  var token = header.slice(start, end);
  if (extensionName === undefined) {
    push(offers, token, params);
  } else {
    if (paramName === undefined) {
      push(params, token, true);
    } else if (mustUnescape) {
      push(params, paramName, token.replace(/\\/g, ''));
    } else {
      push(params, paramName, token);
    }
    push(offers, extensionName, params);
  }
  return offers;
}

/**
 * Builds the `Sec-WebSocket-Extensions` header field value.
 *
 * @param {Object} extensions The map of extensions and parameters to format
 * @return {String} A string representing the given object
 * @public
 */
function format$1(extensions) {
  return Object.keys(extensions).map(function (extension) {
    var configurations = extensions[extension];
    if (!Array.isArray(configurations)) configurations = [configurations];
    return configurations.map(function (params) {
      return [extension].concat(Object.keys(params).map(function (k) {
        var values = params[k];
        if (!Array.isArray(values)) values = [values];
        return values.map(function (v) {
          return v === true ? k : "".concat(k, "=").concat(v);
        }).join('; ');
      })).join('; ');
    }).join(', ');
  }).join(', ');
}
var extension = {
  format: format$1,
  parse: parse$2
};

var EventEmitter = require$$0$4;
var https = require$$4;
var http = require$$3;
var net = require$$3$1;
var tls = require$$4$1;
var randomBytes = require$$5.randomBytes,
  createHash = require$$5.createHash;
require$$0$1.Duplex;
  require$$0$1.Readable;
var URL = require$$1.URL;
var PerMessageDeflate = permessageDeflate;
var Receiver = receiver;
var Sender = sender;
var BINARY_TYPES = constants.BINARY_TYPES,
  EMPTY_BUFFER = constants.EMPTY_BUFFER,
  GUID = constants.GUID,
  kForOnEventAttribute = constants.kForOnEventAttribute,
  kListener = constants.kListener,
  kStatusCode = constants.kStatusCode,
  kWebSocket = constants.kWebSocket,
  NOOP = constants.NOOP;
var _require$$12$EventTar = eventTarget.EventTarget,
  addEventListener$1 = _require$$12$EventTar.addEventListener,
  removeEventListener$1 = _require$$12$EventTar.removeEventListener;
var format = extension.format,
  parse$1 = extension.parse;
var toBuffer = bufferUtilExports.toBuffer;
var closeTimeout = 30 * 1000;
var kAborted = Symbol('kAborted');
var protocolVersions = [8, 13];
var readyStates = ['CONNECTING', 'OPEN', 'CLOSING', 'CLOSED'];
var subprotocolRegex = /^[!#$%&'*+\-.0-9A-Z^_`|a-z~]+$/;

/**
 * Class representing a WebSocket.
 *
 * @extends EventEmitter
 */
var WebSocket$1 = /*#__PURE__*/function (_EventEmitter) {
  /**
   * Create a new `WebSocket`.
   *
   * @param {(String|URL)} address The URL to which to connect
   * @param {(String|String[])} [protocols] The subprotocols
   * @param {Object} [options] Connection options
   */
  function WebSocket(address, protocols, options) {
    var _this;
    _classCallCheck(this, WebSocket);
    _this = _callSuper(this, WebSocket);
    _this._binaryType = BINARY_TYPES[0];
    _this._closeCode = 1006;
    _this._closeFrameReceived = false;
    _this._closeFrameSent = false;
    _this._closeMessage = EMPTY_BUFFER;
    _this._closeTimer = null;
    _this._extensions = {};
    _this._paused = false;
    _this._protocol = '';
    _this._readyState = WebSocket.CONNECTING;
    _this._receiver = null;
    _this._sender = null;
    _this._socket = null;
    if (address !== null) {
      _this._bufferedAmount = 0;
      _this._isServer = false;
      _this._redirects = 0;
      if (protocols === undefined) {
        protocols = [];
      } else if (!Array.isArray(protocols)) {
        if (_typeof(protocols) === 'object' && protocols !== null) {
          options = protocols;
          protocols = [];
        } else {
          protocols = [protocols];
        }
      }
      initAsClient(_this, address, protocols, options);
    } else {
      _this._autoPong = options.autoPong;
      _this._isServer = true;
    }
    return _this;
  }

  /**
   * This deviates from the WHATWG interface since ws doesn't support the
   * required default "blob" type (instead we define a custom "nodebuffer"
   * type).
   *
   * @type {String}
   */
  _inherits(WebSocket, _EventEmitter);
  return _createClass(WebSocket, [{
    key: "binaryType",
    get: function get() {
      return this._binaryType;
    },
    set: function set(type) {
      if (!BINARY_TYPES.includes(type)) return;
      this._binaryType = type;

      //
      // Allow to change `binaryType` on the fly.
      //
      if (this._receiver) this._receiver._binaryType = type;
    }

    /**
     * @type {Number}
     */
  }, {
    key: "bufferedAmount",
    get: function get() {
      if (!this._socket) return this._bufferedAmount;
      return this._socket._writableState.length + this._sender._bufferedBytes;
    }

    /**
     * @type {String}
     */
  }, {
    key: "extensions",
    get: function get() {
      return Object.keys(this._extensions).join();
    }

    /**
     * @type {Boolean}
     */
  }, {
    key: "isPaused",
    get: function get() {
      return this._paused;
    }

    /**
     * @type {Function}
     */
    /* istanbul ignore next */
  }, {
    key: "onclose",
    get: function get() {
      return null;
    }

    /**
     * @type {Function}
     */
    /* istanbul ignore next */
  }, {
    key: "onerror",
    get: function get() {
      return null;
    }

    /**
     * @type {Function}
     */
    /* istanbul ignore next */
  }, {
    key: "onopen",
    get: function get() {
      return null;
    }

    /**
     * @type {Function}
     */
    /* istanbul ignore next */
  }, {
    key: "onmessage",
    get: function get() {
      return null;
    }

    /**
     * @type {String}
     */
  }, {
    key: "protocol",
    get: function get() {
      return this._protocol;
    }

    /**
     * @type {Number}
     */
  }, {
    key: "readyState",
    get: function get() {
      return this._readyState;
    }

    /**
     * @type {String}
     */
  }, {
    key: "url",
    get: function get() {
      return this._url;
    }

    /**
     * Set up the socket and the internal resources.
     *
     * @param {Duplex} socket The network socket between the server and client
     * @param {Buffer} head The first packet of the upgraded stream
     * @param {Object} options Options object
     * @param {Boolean} [options.allowSynchronousEvents=false] Specifies whether
     *     any of the `'message'`, `'ping'`, and `'pong'` events can be emitted
     *     multiple times in the same tick
     * @param {Function} [options.generateMask] The function used to generate the
     *     masking key
     * @param {Number} [options.maxPayload=0] The maximum allowed message size
     * @param {Boolean} [options.skipUTF8Validation=false] Specifies whether or
     *     not to skip UTF-8 validation for text and close messages
     * @private
     */
  }, {
    key: "setSocket",
    value: function setSocket(socket, head, options) {
      var receiver = new Receiver({
        allowSynchronousEvents: options.allowSynchronousEvents,
        binaryType: this.binaryType,
        extensions: this._extensions,
        isServer: this._isServer,
        maxPayload: options.maxPayload,
        skipUTF8Validation: options.skipUTF8Validation
      });
      this._sender = new Sender(socket, this._extensions, options.generateMask);
      this._receiver = receiver;
      this._socket = socket;
      receiver[kWebSocket] = this;
      socket[kWebSocket] = this;
      receiver.on('conclude', receiverOnConclude);
      receiver.on('drain', receiverOnDrain);
      receiver.on('error', receiverOnError);
      receiver.on('message', receiverOnMessage);
      receiver.on('ping', receiverOnPing);
      receiver.on('pong', receiverOnPong);

      //
      // These methods may not be available if `socket` is just a `Duplex`.
      //
      if (socket.setTimeout) socket.setTimeout(0);
      if (socket.setNoDelay) socket.setNoDelay();
      if (head.length > 0) socket.unshift(head);
      socket.on('close', socketOnClose);
      socket.on('data', socketOnData);
      socket.on('end', socketOnEnd);
      socket.on('error', socketOnError);
      this._readyState = WebSocket.OPEN;
      this.emit('open');
    }

    /**
     * Emit the `'close'` event.
     *
     * @private
     */
  }, {
    key: "emitClose",
    value: function emitClose() {
      if (!this._socket) {
        this._readyState = WebSocket.CLOSED;
        this.emit('close', this._closeCode, this._closeMessage);
        return;
      }
      if (this._extensions[PerMessageDeflate.extensionName]) {
        this._extensions[PerMessageDeflate.extensionName].cleanup();
      }
      this._receiver.removeAllListeners();
      this._readyState = WebSocket.CLOSED;
      this.emit('close', this._closeCode, this._closeMessage);
    }

    /**
     * Start a closing handshake.
     *
     *          +----------+   +-----------+   +----------+
     *     - - -|ws.close()|-->|close frame|-->|ws.close()|- - -
     *    |     +----------+   +-----------+   +----------+     |
     *          +----------+   +-----------+         |
     * CLOSING  |ws.close()|<--|close frame|<--+-----+       CLOSING
     *          +----------+   +-----------+   |
     *    |           |                        |   +---+        |
     *                +------------------------+-->|fin| - - - -
     *    |         +---+                      |   +---+
     *     - - - - -|fin|<---------------------+
     *              +---+
     *
     * @param {Number} [code] Status code explaining why the connection is closing
     * @param {(String|Buffer)} [data] The reason why the connection is
     *     closing
     * @public
     */
  }, {
    key: "close",
    value: function close(code, data) {
      var _this2 = this;
      if (this.readyState === WebSocket.CLOSED) return;
      if (this.readyState === WebSocket.CONNECTING) {
        var msg = 'WebSocket was closed before the connection was established';
        abortHandshake(this, this._req, msg);
        return;
      }
      if (this.readyState === WebSocket.CLOSING) {
        if (this._closeFrameSent && (this._closeFrameReceived || this._receiver._writableState.errorEmitted)) {
          this._socket.end();
        }
        return;
      }
      this._readyState = WebSocket.CLOSING;
      this._sender.close(code, data, !this._isServer, function (err) {
        //
        // This error is handled by the `'error'` listener on the socket. We only
        // want to know if the close frame has been sent here.
        //
        if (err) return;
        _this2._closeFrameSent = true;
        if (_this2._closeFrameReceived || _this2._receiver._writableState.errorEmitted) {
          _this2._socket.end();
        }
      });

      //
      // Specify a timeout for the closing handshake to complete.
      //
      this._closeTimer = setTimeout(this._socket.destroy.bind(this._socket), closeTimeout);
    }

    /**
     * Pause the socket.
     *
     * @public
     */
  }, {
    key: "pause",
    value: function pause() {
      if (this.readyState === WebSocket.CONNECTING || this.readyState === WebSocket.CLOSED) {
        return;
      }
      this._paused = true;
      this._socket.pause();
    }

    /**
     * Send a ping.
     *
     * @param {*} [data] The data to send
     * @param {Boolean} [mask] Indicates whether or not to mask `data`
     * @param {Function} [cb] Callback which is executed when the ping is sent
     * @public
     */
  }, {
    key: "ping",
    value: function ping(data, mask, cb) {
      if (this.readyState === WebSocket.CONNECTING) {
        throw new Error('WebSocket is not open: readyState 0 (CONNECTING)');
      }
      if (typeof data === 'function') {
        cb = data;
        data = mask = undefined;
      } else if (typeof mask === 'function') {
        cb = mask;
        mask = undefined;
      }
      if (typeof data === 'number') data = data.toString();
      if (this.readyState !== WebSocket.OPEN) {
        sendAfterClose(this, data, cb);
        return;
      }
      if (mask === undefined) mask = !this._isServer;
      this._sender.ping(data || EMPTY_BUFFER, mask, cb);
    }

    /**
     * Send a pong.
     *
     * @param {*} [data] The data to send
     * @param {Boolean} [mask] Indicates whether or not to mask `data`
     * @param {Function} [cb] Callback which is executed when the pong is sent
     * @public
     */
  }, {
    key: "pong",
    value: function pong(data, mask, cb) {
      if (this.readyState === WebSocket.CONNECTING) {
        throw new Error('WebSocket is not open: readyState 0 (CONNECTING)');
      }
      if (typeof data === 'function') {
        cb = data;
        data = mask = undefined;
      } else if (typeof mask === 'function') {
        cb = mask;
        mask = undefined;
      }
      if (typeof data === 'number') data = data.toString();
      if (this.readyState !== WebSocket.OPEN) {
        sendAfterClose(this, data, cb);
        return;
      }
      if (mask === undefined) mask = !this._isServer;
      this._sender.pong(data || EMPTY_BUFFER, mask, cb);
    }

    /**
     * Resume the socket.
     *
     * @public
     */
  }, {
    key: "resume",
    value: function resume() {
      if (this.readyState === WebSocket.CONNECTING || this.readyState === WebSocket.CLOSED) {
        return;
      }
      this._paused = false;
      if (!this._receiver._writableState.needDrain) this._socket.resume();
    }

    /**
     * Send a data message.
     *
     * @param {*} data The message to send
     * @param {Object} [options] Options object
     * @param {Boolean} [options.binary] Specifies whether `data` is binary or
     *     text
     * @param {Boolean} [options.compress] Specifies whether or not to compress
     *     `data`
     * @param {Boolean} [options.fin=true] Specifies whether the fragment is the
     *     last one
     * @param {Boolean} [options.mask] Specifies whether or not to mask `data`
     * @param {Function} [cb] Callback which is executed when data is written out
     * @public
     */
  }, {
    key: "send",
    value: function send(data, options, cb) {
      if (this.readyState === WebSocket.CONNECTING) {
        throw new Error('WebSocket is not open: readyState 0 (CONNECTING)');
      }
      if (typeof options === 'function') {
        cb = options;
        options = {};
      }
      if (typeof data === 'number') data = data.toString();
      if (this.readyState !== WebSocket.OPEN) {
        sendAfterClose(this, data, cb);
        return;
      }
      var opts = _objectSpread2({
        binary: typeof data !== 'string',
        mask: !this._isServer,
        compress: true,
        fin: true
      }, options);
      if (!this._extensions[PerMessageDeflate.extensionName]) {
        opts.compress = false;
      }
      this._sender.send(data || EMPTY_BUFFER, opts, cb);
    }

    /**
     * Forcibly close the connection.
     *
     * @public
     */
  }, {
    key: "terminate",
    value: function terminate() {
      if (this.readyState === WebSocket.CLOSED) return;
      if (this.readyState === WebSocket.CONNECTING) {
        var msg = 'WebSocket was closed before the connection was established';
        abortHandshake(this, this._req, msg);
        return;
      }
      if (this._socket) {
        this._readyState = WebSocket.CLOSING;
        this._socket.destroy();
      }
    }
  }]);
}(EventEmitter);
/**
 * @constant {Number} CONNECTING
 * @memberof WebSocket
 */
Object.defineProperty(WebSocket$1, 'CONNECTING', {
  enumerable: true,
  value: readyStates.indexOf('CONNECTING')
});

/**
 * @constant {Number} CONNECTING
 * @memberof WebSocket.prototype
 */
Object.defineProperty(WebSocket$1.prototype, 'CONNECTING', {
  enumerable: true,
  value: readyStates.indexOf('CONNECTING')
});

/**
 * @constant {Number} OPEN
 * @memberof WebSocket
 */
Object.defineProperty(WebSocket$1, 'OPEN', {
  enumerable: true,
  value: readyStates.indexOf('OPEN')
});

/**
 * @constant {Number} OPEN
 * @memberof WebSocket.prototype
 */
Object.defineProperty(WebSocket$1.prototype, 'OPEN', {
  enumerable: true,
  value: readyStates.indexOf('OPEN')
});

/**
 * @constant {Number} CLOSING
 * @memberof WebSocket
 */
Object.defineProperty(WebSocket$1, 'CLOSING', {
  enumerable: true,
  value: readyStates.indexOf('CLOSING')
});

/**
 * @constant {Number} CLOSING
 * @memberof WebSocket.prototype
 */
Object.defineProperty(WebSocket$1.prototype, 'CLOSING', {
  enumerable: true,
  value: readyStates.indexOf('CLOSING')
});

/**
 * @constant {Number} CLOSED
 * @memberof WebSocket
 */
Object.defineProperty(WebSocket$1, 'CLOSED', {
  enumerable: true,
  value: readyStates.indexOf('CLOSED')
});

/**
 * @constant {Number} CLOSED
 * @memberof WebSocket.prototype
 */
Object.defineProperty(WebSocket$1.prototype, 'CLOSED', {
  enumerable: true,
  value: readyStates.indexOf('CLOSED')
});
['binaryType', 'bufferedAmount', 'extensions', 'isPaused', 'protocol', 'readyState', 'url'].forEach(function (property) {
  Object.defineProperty(WebSocket$1.prototype, property, {
    enumerable: true
  });
});

//
// Add the `onopen`, `onerror`, `onclose`, and `onmessage` attributes.
// See https://html.spec.whatwg.org/multipage/comms.html#the-websocket-interface
//
['open', 'error', 'close', 'message'].forEach(function (method) {
  Object.defineProperty(WebSocket$1.prototype, "on".concat(method), {
    enumerable: true,
    get: function get() {
      var _iterator = _createForOfIteratorHelper(this.listeners(method)),
        _step;
      try {
        for (_iterator.s(); !(_step = _iterator.n()).done;) {
          var listener = _step.value;
          if (listener[kForOnEventAttribute]) return listener[kListener];
        }
      } catch (err) {
        _iterator.e(err);
      } finally {
        _iterator.f();
      }
      return null;
    },
    set: function set(handler) {
      var _iterator2 = _createForOfIteratorHelper(this.listeners(method)),
        _step2;
      try {
        for (_iterator2.s(); !(_step2 = _iterator2.n()).done;) {
          var listener = _step2.value;
          if (listener[kForOnEventAttribute]) {
            this.removeListener(method, listener);
            break;
          }
        }
      } catch (err) {
        _iterator2.e(err);
      } finally {
        _iterator2.f();
      }
      if (typeof handler !== 'function') return;
      this.addEventListener(method, handler, _defineProperty({}, kForOnEventAttribute, true));
    }
  });
});
WebSocket$1.prototype.addEventListener = addEventListener$1;
WebSocket$1.prototype.removeEventListener = removeEventListener$1;
var websocket = WebSocket$1;

/**
 * Initialize a WebSocket client.
 *
 * @param {WebSocket} websocket The client to initialize
 * @param {(String|URL)} address The URL to which to connect
 * @param {Array} protocols The subprotocols
 * @param {Object} [options] Connection options
 * @param {Boolean} [options.allowSynchronousEvents=true] Specifies whether any
 *     of the `'message'`, `'ping'`, and `'pong'` events can be emitted multiple
 *     times in the same tick
 * @param {Boolean} [options.autoPong=true] Specifies whether or not to
 *     automatically send a pong in response to a ping
 * @param {Function} [options.finishRequest] A function which can be used to
 *     customize the headers of each http request before it is sent
 * @param {Boolean} [options.followRedirects=false] Whether or not to follow
 *     redirects
 * @param {Function} [options.generateMask] The function used to generate the
 *     masking key
 * @param {Number} [options.handshakeTimeout] Timeout in milliseconds for the
 *     handshake request
 * @param {Number} [options.maxPayload=104857600] The maximum allowed message
 *     size
 * @param {Number} [options.maxRedirects=10] The maximum number of redirects
 *     allowed
 * @param {String} [options.origin] Value of the `Origin` or
 *     `Sec-WebSocket-Origin` header
 * @param {(Boolean|Object)} [options.perMessageDeflate=true] Enable/disable
 *     permessage-deflate
 * @param {Number} [options.protocolVersion=13] Value of the
 *     `Sec-WebSocket-Version` header
 * @param {Boolean} [options.skipUTF8Validation=false] Specifies whether or
 *     not to skip UTF-8 validation for text and close messages
 * @private
 */
function initAsClient(websocket, address, protocols, options) {
  var opts = _objectSpread2(_objectSpread2({
    allowSynchronousEvents: true,
    autoPong: true,
    protocolVersion: protocolVersions[1],
    maxPayload: 100 * 1024 * 1024,
    skipUTF8Validation: false,
    perMessageDeflate: true,
    followRedirects: false,
    maxRedirects: 10
  }, options), {}, {
    socketPath: undefined,
    hostname: undefined,
    protocol: undefined,
    timeout: undefined,
    method: 'GET',
    host: undefined,
    path: undefined,
    port: undefined
  });
  websocket._autoPong = opts.autoPong;
  if (!protocolVersions.includes(opts.protocolVersion)) {
    throw new RangeError("Unsupported protocol version: ".concat(opts.protocolVersion, " ") + "(supported versions: ".concat(protocolVersions.join(', '), ")"));
  }
  var parsedUrl;
  if (address instanceof URL) {
    parsedUrl = address;
  } else {
    try {
      parsedUrl = new URL(address);
    } catch (e) {
      throw new SyntaxError("Invalid URL: ".concat(address));
    }
  }
  if (parsedUrl.protocol === 'http:') {
    parsedUrl.protocol = 'ws:';
  } else if (parsedUrl.protocol === 'https:') {
    parsedUrl.protocol = 'wss:';
  }
  websocket._url = parsedUrl.href;
  var isSecure = parsedUrl.protocol === 'wss:';
  var isIpcUrl = parsedUrl.protocol === 'ws+unix:';
  var invalidUrlMessage;
  if (parsedUrl.protocol !== 'ws:' && !isSecure && !isIpcUrl) {
    invalidUrlMessage = 'The URL\'s protocol must be one of "ws:", "wss:", ' + '"http:", "https", or "ws+unix:"';
  } else if (isIpcUrl && !parsedUrl.pathname) {
    invalidUrlMessage = "The URL's pathname is empty";
  } else if (parsedUrl.hash) {
    invalidUrlMessage = 'The URL contains a fragment identifier';
  }
  if (invalidUrlMessage) {
    var err = new SyntaxError(invalidUrlMessage);
    if (websocket._redirects === 0) {
      throw err;
    } else {
      emitErrorAndClose(websocket, err);
      return;
    }
  }
  var defaultPort = isSecure ? 443 : 80;
  var key = randomBytes(16).toString('base64');
  var request = isSecure ? https.request : http.request;
  var protocolSet = new Set();
  var perMessageDeflate;
  opts.createConnection = opts.createConnection || (isSecure ? tlsConnect : netConnect);
  opts.defaultPort = opts.defaultPort || defaultPort;
  opts.port = parsedUrl.port || defaultPort;
  opts.host = parsedUrl.hostname.startsWith('[') ? parsedUrl.hostname.slice(1, -1) : parsedUrl.hostname;
  opts.headers = _objectSpread2(_objectSpread2({}, opts.headers), {}, {
    'Sec-WebSocket-Version': opts.protocolVersion,
    'Sec-WebSocket-Key': key,
    Connection: 'Upgrade',
    Upgrade: 'websocket'
  });
  opts.path = parsedUrl.pathname + parsedUrl.search;
  opts.timeout = opts.handshakeTimeout;
  if (opts.perMessageDeflate) {
    perMessageDeflate = new PerMessageDeflate(opts.perMessageDeflate !== true ? opts.perMessageDeflate : {}, false, opts.maxPayload);
    opts.headers['Sec-WebSocket-Extensions'] = format(_defineProperty({}, PerMessageDeflate.extensionName, perMessageDeflate.offer()));
  }
  if (protocols.length) {
    var _iterator3 = _createForOfIteratorHelper(protocols),
      _step3;
    try {
      for (_iterator3.s(); !(_step3 = _iterator3.n()).done;) {
        var protocol = _step3.value;
        if (typeof protocol !== 'string' || !subprotocolRegex.test(protocol) || protocolSet.has(protocol)) {
          throw new SyntaxError('An invalid or duplicated subprotocol was specified');
        }
        protocolSet.add(protocol);
      }
    } catch (err) {
      _iterator3.e(err);
    } finally {
      _iterator3.f();
    }
    opts.headers['Sec-WebSocket-Protocol'] = protocols.join(',');
  }
  if (opts.origin) {
    if (opts.protocolVersion < 13) {
      opts.headers['Sec-WebSocket-Origin'] = opts.origin;
    } else {
      opts.headers.Origin = opts.origin;
    }
  }
  if (parsedUrl.username || parsedUrl.password) {
    opts.auth = "".concat(parsedUrl.username, ":").concat(parsedUrl.password);
  }
  if (isIpcUrl) {
    var parts = opts.path.split(':');
    opts.socketPath = parts[0];
    opts.path = parts[1];
  }
  var req;
  if (opts.followRedirects) {
    if (websocket._redirects === 0) {
      websocket._originalIpc = isIpcUrl;
      websocket._originalSecure = isSecure;
      websocket._originalHostOrSocketPath = isIpcUrl ? opts.socketPath : parsedUrl.host;
      var headers = options && options.headers;

      //
      // Shallow copy the user provided options so that headers can be changed
      // without mutating the original object.
      //
      options = _objectSpread2(_objectSpread2({}, options), {}, {
        headers: {}
      });
      if (headers) {
        for (var _i = 0, _Object$entries = Object.entries(headers); _i < _Object$entries.length; _i++) {
          var _Object$entries$_i = _slicedToArray(_Object$entries[_i], 2),
            _key = _Object$entries$_i[0],
            value = _Object$entries$_i[1];
          options.headers[_key.toLowerCase()] = value;
        }
      }
    } else if (websocket.listenerCount('redirect') === 0) {
      var isSameHost = isIpcUrl ? websocket._originalIpc ? opts.socketPath === websocket._originalHostOrSocketPath : false : websocket._originalIpc ? false : parsedUrl.host === websocket._originalHostOrSocketPath;
      if (!isSameHost || websocket._originalSecure && !isSecure) {
        //
        // Match curl 7.77.0 behavior and drop the following headers. These
        // headers are also dropped when following a redirect to a subdomain.
        //
        delete opts.headers.authorization;
        delete opts.headers.cookie;
        if (!isSameHost) delete opts.headers.host;
        opts.auth = undefined;
      }
    }

    //
    // Match curl 7.77.0 behavior and make the first `Authorization` header win.
    // If the `Authorization` header is set, then there is nothing to do as it
    // will take precedence.
    //
    if (opts.auth && !options.headers.authorization) {
      options.headers.authorization = 'Basic ' + Buffer.from(opts.auth).toString('base64');
    }
    req = websocket._req = request(opts);
    if (websocket._redirects) {
      //
      // Unlike what is done for the `'upgrade'` event, no early exit is
      // triggered here if the user calls `websocket.close()` or
      // `websocket.terminate()` from a listener of the `'redirect'` event. This
      // is because the user can also call `request.destroy()` with an error
      // before calling `websocket.close()` or `websocket.terminate()` and this
      // would result in an error being emitted on the `request` object with no
      // `'error'` event listeners attached.
      //
      websocket.emit('redirect', websocket.url, req);
    }
  } else {
    req = websocket._req = request(opts);
  }
  if (opts.timeout) {
    req.on('timeout', function () {
      abortHandshake(websocket, req, 'Opening handshake has timed out');
    });
  }
  req.on('error', function (err) {
    if (req === null || req[kAborted]) return;
    req = websocket._req = null;
    emitErrorAndClose(websocket, err);
  });
  req.on('response', function (res) {
    var location = res.headers.location;
    var statusCode = res.statusCode;
    if (location && opts.followRedirects && statusCode >= 300 && statusCode < 400) {
      if (++websocket._redirects > opts.maxRedirects) {
        abortHandshake(websocket, req, 'Maximum redirects exceeded');
        return;
      }
      req.abort();
      var addr;
      try {
        addr = new URL(location, address);
      } catch (e) {
        var _err = new SyntaxError("Invalid URL: ".concat(location));
        emitErrorAndClose(websocket, _err);
        return;
      }
      initAsClient(websocket, addr, protocols, options);
    } else if (!websocket.emit('unexpected-response', req, res)) {
      abortHandshake(websocket, req, "Unexpected server response: ".concat(res.statusCode));
    }
  });
  req.on('upgrade', function (res, socket, head) {
    websocket.emit('upgrade', res);

    //
    // The user may have closed the connection from a listener of the
    // `'upgrade'` event.
    //
    if (websocket.readyState !== WebSocket$1.CONNECTING) return;
    req = websocket._req = null;
    var upgrade = res.headers.upgrade;
    if (upgrade === undefined || upgrade.toLowerCase() !== 'websocket') {
      abortHandshake(websocket, socket, 'Invalid Upgrade header');
      return;
    }
    var digest = createHash('sha1').update(key + GUID).digest('base64');
    if (res.headers['sec-websocket-accept'] !== digest) {
      abortHandshake(websocket, socket, 'Invalid Sec-WebSocket-Accept header');
      return;
    }
    var serverProt = res.headers['sec-websocket-protocol'];
    var protError;
    if (serverProt !== undefined) {
      if (!protocolSet.size) {
        protError = 'Server sent a subprotocol but none was requested';
      } else if (!protocolSet.has(serverProt)) {
        protError = 'Server sent an invalid subprotocol';
      }
    } else if (protocolSet.size) {
      protError = 'Server sent no subprotocol';
    }
    if (protError) {
      abortHandshake(websocket, socket, protError);
      return;
    }
    if (serverProt) websocket._protocol = serverProt;
    var secWebSocketExtensions = res.headers['sec-websocket-extensions'];
    if (secWebSocketExtensions !== undefined) {
      if (!perMessageDeflate) {
        var message = 'Server sent a Sec-WebSocket-Extensions header but no extension ' + 'was requested';
        abortHandshake(websocket, socket, message);
        return;
      }
      var extensions;
      try {
        extensions = parse$1(secWebSocketExtensions);
      } catch (err) {
        var _message = 'Invalid Sec-WebSocket-Extensions header';
        abortHandshake(websocket, socket, _message);
        return;
      }
      var extensionNames = Object.keys(extensions);
      if (extensionNames.length !== 1 || extensionNames[0] !== PerMessageDeflate.extensionName) {
        var _message2 = 'Server indicated an extension that was not requested';
        abortHandshake(websocket, socket, _message2);
        return;
      }
      try {
        perMessageDeflate.accept(extensions[PerMessageDeflate.extensionName]);
      } catch (err) {
        var _message3 = 'Invalid Sec-WebSocket-Extensions header';
        abortHandshake(websocket, socket, _message3);
        return;
      }
      websocket._extensions[PerMessageDeflate.extensionName] = perMessageDeflate;
    }
    websocket.setSocket(socket, head, {
      allowSynchronousEvents: opts.allowSynchronousEvents,
      generateMask: opts.generateMask,
      maxPayload: opts.maxPayload,
      skipUTF8Validation: opts.skipUTF8Validation
    });
  });
  if (opts.finishRequest) {
    opts.finishRequest(req, websocket);
  } else {
    req.end();
  }
}

/**
 * Emit the `'error'` and `'close'` events.
 *
 * @param {WebSocket} websocket The WebSocket instance
 * @param {Error} The error to emit
 * @private
 */
function emitErrorAndClose(websocket, err) {
  websocket._readyState = WebSocket$1.CLOSING;
  websocket.emit('error', err);
  websocket.emitClose();
}

/**
 * Create a `net.Socket` and initiate a connection.
 *
 * @param {Object} options Connection options
 * @return {net.Socket} The newly created socket used to start the connection
 * @private
 */
function netConnect(options) {
  options.path = options.socketPath;
  return net.connect(options);
}

/**
 * Create a `tls.TLSSocket` and initiate a connection.
 *
 * @param {Object} options Connection options
 * @return {tls.TLSSocket} The newly created socket used to start the connection
 * @private
 */
function tlsConnect(options) {
  options.path = undefined;
  if (!options.servername && options.servername !== '') {
    options.servername = net.isIP(options.host) ? '' : options.host;
  }
  return tls.connect(options);
}

/**
 * Abort the handshake and emit an error.
 *
 * @param {WebSocket} websocket The WebSocket instance
 * @param {(http.ClientRequest|net.Socket|tls.Socket)} stream The request to
 *     abort or the socket to destroy
 * @param {String} message The error message
 * @private
 */
function abortHandshake(websocket, stream, message) {
  websocket._readyState = WebSocket$1.CLOSING;
  var err = new Error(message);
  Error.captureStackTrace(err, abortHandshake);
  if (stream.setHeader) {
    stream[kAborted] = true;
    stream.abort();
    if (stream.socket && !stream.socket.destroyed) {
      //
      // On Node.js >= 14.3.0 `request.abort()` does not destroy the socket if
      // called after the request completed. See
      // https://github.com/websockets/ws/issues/1869.
      //
      stream.socket.destroy();
    }
    process.nextTick(emitErrorAndClose, websocket, err);
  } else {
    stream.destroy(err);
    stream.once('error', websocket.emit.bind(websocket, 'error'));
    stream.once('close', websocket.emitClose.bind(websocket));
  }
}

/**
 * Handle cases where the `ping()`, `pong()`, or `send()` methods are called
 * when the `readyState` attribute is `CLOSING` or `CLOSED`.
 *
 * @param {WebSocket} websocket The WebSocket instance
 * @param {*} [data] The data to send
 * @param {Function} [cb] Callback
 * @private
 */
function sendAfterClose(websocket, data, cb) {
  if (data) {
    var length = toBuffer(data).length;

    //
    // The `_bufferedAmount` property is used only when the peer is a client and
    // the opening handshake fails. Under these circumstances, in fact, the
    // `setSocket()` method is not called, so the `_socket` and `_sender`
    // properties are set to `null`.
    //
    if (websocket._socket) websocket._sender._bufferedBytes += length;else websocket._bufferedAmount += length;
  }
  if (cb) {
    var err = new Error("WebSocket is not open: readyState ".concat(websocket.readyState, " ") + "(".concat(readyStates[websocket.readyState], ")"));
    process.nextTick(cb, err);
  }
}

/**
 * The listener of the `Receiver` `'conclude'` event.
 *
 * @param {Number} code The status code
 * @param {Buffer} reason The reason for closing
 * @private
 */
function receiverOnConclude(code, reason) {
  var websocket = this[kWebSocket];
  websocket._closeFrameReceived = true;
  websocket._closeMessage = reason;
  websocket._closeCode = code;
  if (websocket._socket[kWebSocket] === undefined) return;
  websocket._socket.removeListener('data', socketOnData);
  process.nextTick(resume, websocket._socket);
  if (code === 1005) websocket.close();else websocket.close(code, reason);
}

/**
 * The listener of the `Receiver` `'drain'` event.
 *
 * @private
 */
function receiverOnDrain() {
  var websocket = this[kWebSocket];
  if (!websocket.isPaused) websocket._socket.resume();
}

/**
 * The listener of the `Receiver` `'error'` event.
 *
 * @param {(RangeError|Error)} err The emitted error
 * @private
 */
function receiverOnError(err) {
  var websocket = this[kWebSocket];
  if (websocket._socket[kWebSocket] !== undefined) {
    websocket._socket.removeListener('data', socketOnData);

    //
    // On Node.js < 14.0.0 the `'error'` event is emitted synchronously. See
    // https://github.com/websockets/ws/issues/1940.
    //
    process.nextTick(resume, websocket._socket);
    websocket.close(err[kStatusCode]);
  }
  websocket.emit('error', err);
}

/**
 * The listener of the `Receiver` `'finish'` event.
 *
 * @private
 */
function receiverOnFinish() {
  this[kWebSocket].emitClose();
}

/**
 * The listener of the `Receiver` `'message'` event.
 *
 * @param {Buffer|ArrayBuffer|Buffer[])} data The message
 * @param {Boolean} isBinary Specifies whether the message is binary or not
 * @private
 */
function receiverOnMessage(data, isBinary) {
  this[kWebSocket].emit('message', data, isBinary);
}

/**
 * The listener of the `Receiver` `'ping'` event.
 *
 * @param {Buffer} data The data included in the ping frame
 * @private
 */
function receiverOnPing(data) {
  var websocket = this[kWebSocket];
  if (websocket._autoPong) websocket.pong(data, !this._isServer, NOOP);
  websocket.emit('ping', data);
}

/**
 * The listener of the `Receiver` `'pong'` event.
 *
 * @param {Buffer} data The data included in the pong frame
 * @private
 */
function receiverOnPong(data) {
  this[kWebSocket].emit('pong', data);
}

/**
 * Resume a readable stream
 *
 * @param {Readable} stream The readable stream
 * @private
 */
function resume(stream) {
  stream.resume();
}

/**
 * The listener of the socket `'close'` event.
 *
 * @private
 */
function socketOnClose() {
  var websocket = this[kWebSocket];
  this.removeListener('close', socketOnClose);
  this.removeListener('data', socketOnData);
  this.removeListener('end', socketOnEnd);
  websocket._readyState = WebSocket$1.CLOSING;
  var chunk;

  //
  // The close frame might not have been received or the `'end'` event emitted,
  // for example, if the socket was destroyed due to an error. Ensure that the
  // `receiver` stream is closed after writing any remaining buffered data to
  // it. If the readable side of the socket is in flowing mode then there is no
  // buffered data as everything has been already written and `readable.read()`
  // will return `null`. If instead, the socket is paused, any possible buffered
  // data will be read as a single chunk.
  //
  if (!this._readableState.endEmitted && !websocket._closeFrameReceived && !websocket._receiver._writableState.errorEmitted && (chunk = websocket._socket.read()) !== null) {
    websocket._receiver.write(chunk);
  }
  websocket._receiver.end();
  this[kWebSocket] = undefined;
  clearTimeout(websocket._closeTimer);
  if (websocket._receiver._writableState.finished || websocket._receiver._writableState.errorEmitted) {
    websocket.emitClose();
  } else {
    websocket._receiver.on('error', receiverOnFinish);
    websocket._receiver.on('finish', receiverOnFinish);
  }
}

/**
 * The listener of the socket `'data'` event.
 *
 * @param {Buffer} chunk A chunk of data
 * @private
 */
function socketOnData(chunk) {
  if (!this[kWebSocket]._receiver.write(chunk)) {
    this.pause();
  }
}

/**
 * The listener of the socket `'end'` event.
 *
 * @private
 */
function socketOnEnd() {
  var websocket = this[kWebSocket];
  websocket._readyState = WebSocket$1.CLOSING;
  websocket._receiver.end();
  this.end();
}

/**
 * The listener of the socket `'error'` event.
 *
 * @private
 */
function socketOnError() {
  var websocket = this[kWebSocket];
  this.removeListener('error', socketOnError);
  this.on('error', NOOP);
  if (websocket) {
    websocket._readyState = WebSocket$1.CLOSING;
    this.destroy();
  }
}
var WebSocket$2 = /*@__PURE__*/getDefaultExportFromCjs(websocket);

validationExports.tokenChars;

require$$0$1.Duplex;
require$$5.createHash;
constants.GUID;
  constants.kWebSocket;

var WebSocket = WebSocket$2;
var usingBrowserWebSocket = false;
var defaultBinaryType = "nodebuffer";
var nextTick = process.nextTick;

// detect ReactNative environment
var isReactNative = typeof navigator !== "undefined" && typeof navigator.product === "string" && navigator.product.toLowerCase() === "reactnative";
var WS = /*#__PURE__*/function (_Transport) {
  /**
   * WebSocket transport constructor.
   *
   * @param {Object} opts - connection options
   * @protected
   */
  function WS(opts) {
    var _this;
    _classCallCheck(this, WS);
    _this = _callSuper(this, WS, [opts]);
    _this.supportsBinary = !opts.forceBase64;
    return _this;
  }
  _inherits(WS, _Transport);
  return _createClass(WS, [{
    key: "name",
    get: function get() {
      return "websocket";
    }
  }, {
    key: "doOpen",
    value: function doOpen() {
      if (!this.check()) {
        // let probe timeout
        return;
      }
      var uri = this.uri();
      var protocols = this.opts.protocols;
      // React Native only supports the 'headers' option, and will print a warning if anything else is passed
      var opts = isReactNative ? {} : pick(this.opts, "agent", "perMessageDeflate", "pfx", "key", "passphrase", "cert", "ca", "ciphers", "rejectUnauthorized", "localAddress", "protocolVersion", "origin", "maxPayload", "family", "checkServerIdentity");
      if (this.opts.extraHeaders) {
        opts.headers = this.opts.extraHeaders;
      }
      try {
        this.ws = usingBrowserWebSocket && !isReactNative ? protocols ? new WebSocket(uri, protocols) : new WebSocket(uri) : new WebSocket(uri, protocols, opts);
      } catch (err) {
        return this.emitReserved("error", err);
      }
      this.ws.binaryType = this.socket.binaryType;
      this.addEventListeners();
    }
    /**
     * Adds event listeners to the socket
     *
     * @private
     */
  }, {
    key: "addEventListeners",
    value: function addEventListeners() {
      var _this2 = this;
      this.ws.onopen = function () {
        if (_this2.opts.autoUnref) {
          _this2.ws._socket.unref();
        }
        _this2.onOpen();
      };
      this.ws.onclose = function (closeEvent) {
        return _this2.onClose({
          description: "websocket connection closed",
          context: closeEvent
        });
      };
      this.ws.onmessage = function (ev) {
        return _this2.onData(ev.data);
      };
      this.ws.onerror = function (e) {
        return _this2.onError("websocket error", e);
      };
    }
  }, {
    key: "write",
    value: function write(packets) {
      var _this3 = this;
      this.writable = false;
      // encodePacket efficient as it uses WS framing
      // no need for encodePayload
      var _loop = function _loop() {
        var packet = packets[i];
        var lastPacket = i === packets.length - 1;
        encodePacket(packet, _this3.supportsBinary, function (data) {
          // always create a new object (GH-437)
          var opts = {};
          {
            if (packet.options) {
              opts.compress = packet.options.compress;
            }
            if (_this3.opts.perMessageDeflate) {
              var len =
              // @ts-ignore
              "string" === typeof data ? Buffer.byteLength(data) : data.length;
              if (len < _this3.opts.perMessageDeflate.threshold) {
                opts.compress = false;
              }
            }
          }
          // Sometimes the websocket has already been closed but the browser didn't
          // have a chance of informing us about it yet, in that case send will
          // throw an error
          try {
            if (usingBrowserWebSocket) ; else {
              _this3.ws.send(data, opts);
            }
          } catch (e) {}
          if (lastPacket) {
            // fake drain
            // defer to next tick to allow Socket to clear writeBuffer
            nextTick(function () {
              _this3.writable = true;
              _this3.emitReserved("drain");
            }, _this3.setTimeoutFn);
          }
        });
      };
      for (var i = 0; i < packets.length; i++) {
        _loop();
      }
    }
  }, {
    key: "doClose",
    value: function doClose() {
      if (typeof this.ws !== "undefined") {
        this.ws.close();
        this.ws = null;
      }
    }
    /**
     * Generates uri for connection.
     *
     * @private
     */
  }, {
    key: "uri",
    value: function uri() {
      var schema = this.opts.secure ? "wss" : "ws";
      var query = this.query || {};
      // append timestamp to URI
      if (this.opts.timestampRequests) {
        query[this.opts.timestampParam] = yeast();
      }
      // communicate binary support capabilities
      if (!this.supportsBinary) {
        query.b64 = 1;
      }
      return this.createUri(schema, query);
    }
    /**
     * Feature detection for WebSocket.
     *
     * @return {Boolean} whether this transport is available.
     * @private
     */
  }, {
    key: "check",
    value: function check() {
      return !!WebSocket;
    }
  }]);
}(Transport);

var WT = /*#__PURE__*/function (_Transport) {
  function WT() {
    _classCallCheck(this, WT);
    return _callSuper(this, WT, arguments);
  }
  _inherits(WT, _Transport);
  return _createClass(WT, [{
    key: "name",
    get: function get() {
      return "webtransport";
    }
  }, {
    key: "doOpen",
    value: function doOpen() {
      var _this = this;
      // @ts-ignore
      if (typeof WebTransport !== "function") {
        return;
      }
      // @ts-ignore
      this.transport = new WebTransport(this.createUri("https"), this.opts.transportOptions[this.name]);
      this.transport.closed.then(function () {
        _this.onClose();
      })["catch"](function (err) {
        _this.onError("webtransport error", err);
      });
      // note: we could have used async/await, but that would require some additional polyfills
      this.transport.ready.then(function () {
        _this.transport.createBidirectionalStream().then(function (stream) {
          var decoderStream = createPacketDecoderStream(Number.MAX_SAFE_INTEGER, _this.socket.binaryType);
          var reader = stream.readable.pipeThrough(decoderStream).getReader();
          var encoderStream = createPacketEncoderStream();
          encoderStream.readable.pipeTo(stream.writable);
          _this.writer = encoderStream.writable.getWriter();
          var _read = function read() {
            reader.read().then(function (_ref) {
              var done = _ref.done,
                value = _ref.value;
              if (done) {
                return;
              }
              _this.onPacket(value);
              _read();
            })["catch"](function (err) {});
          };
          _read();
          var packet = {
            type: "open"
          };
          if (_this.query.sid) {
            packet.data = "{\"sid\":\"".concat(_this.query.sid, "\"}");
          }
          _this.writer.write(packet).then(function () {
            return _this.onOpen();
          });
        });
      });
    }
  }, {
    key: "write",
    value: function write(packets) {
      var _this2 = this;
      this.writable = false;
      var _loop = function _loop() {
        var packet = packets[i];
        var lastPacket = i === packets.length - 1;
        _this2.writer.write(packet).then(function () {
          if (lastPacket) {
            nextTick(function () {
              _this2.writable = true;
              _this2.emitReserved("drain");
            }, _this2.setTimeoutFn);
          }
        });
      };
      for (var i = 0; i < packets.length; i++) {
        _loop();
      }
    }
  }, {
    key: "doClose",
    value: function doClose() {
      var _a;
      (_a = this.transport) === null || _a === void 0 ? void 0 : _a.close();
    }
  }]);
}(Transport);

var transports = {
  websocket: WS,
  webtransport: WT,
  polling: Polling
};

// imported from https://github.com/galkn/parseuri
/**
 * Parses a URI
 *
 * Note: we could also have used the built-in URL object, but it isn't supported on all platforms.
 *
 * See:
 * - https://developer.mozilla.org/en-US/docs/Web/API/URL
 * - https://caniuse.com/url
 * - https://www.rfc-editor.org/rfc/rfc3986#appendix-B
 *
 * History of the parse() method:
 * - first commit: https://github.com/socketio/socket.io-client/commit/4ee1d5d94b3906a9c052b459f1a818b15f38f91c
 * - export into its own module: https://github.com/socketio/engine.io-client/commit/de2c561e4564efeb78f1bdb1ba39ef81b2822cb3
 * - reimport: https://github.com/socketio/engine.io-client/commit/df32277c3f6d622eec5ed09f493cae3f3391d242
 *
 * @author Steven Levithan <stevenlevithan.com> (MIT license)
 * @api private
 */
var re = /^(?:(?![^:@\/?#]+:[^:@\/]*@)(http|https|ws|wss):\/\/)?((?:(([^:@\/?#]*)(?::([^:@\/?#]*))?)?@)?((?:[a-f0-9]{0,4}:){2,7}[a-f0-9]{0,4}|[^:\/?#]*)(?::(\d*))?)(((\/(?:[^?#](?![^?#\/]*\.[^?#\/.]+(?:[?#]|$)))*\/?)?([^?#\/]*))(?:\?([^#]*))?(?:#(.*))?)/;
var parts = ['source', 'protocol', 'authority', 'userInfo', 'user', 'password', 'host', 'port', 'relative', 'path', 'directory', 'file', 'query', 'anchor'];
function parse(str) {
  if (str.length > 2000) {
    throw "URI too long";
  }
  var src = str,
    b = str.indexOf('['),
    e = str.indexOf(']');
  if (b != -1 && e != -1) {
    str = str.substring(0, b) + str.substring(b, e).replace(/:/g, ';') + str.substring(e, str.length);
  }
  var m = re.exec(str || ''),
    uri = {},
    i = 14;
  while (i--) {
    uri[parts[i]] = m[i] || '';
  }
  if (b != -1 && e != -1) {
    uri.source = src;
    uri.host = uri.host.substring(1, uri.host.length - 1).replace(/;/g, ':');
    uri.authority = uri.authority.replace('[', '').replace(']', '').replace(/;/g, ':');
    uri.ipv6uri = true;
  }
  uri.pathNames = pathNames(uri, uri['path']);
  uri.queryKey = queryKey(uri, uri['query']);
  return uri;
}
function pathNames(obj, path) {
  var regx = /\/{2,9}/g,
    names = path.replace(regx, "/").split("/");
  if (path.slice(0, 1) == '/' || path.length === 0) {
    names.splice(0, 1);
  }
  if (path.slice(-1) == '/') {
    names.splice(names.length - 1, 1);
  }
  return names;
}
function queryKey(uri, query) {
  var data = {};
  query.replace(/(?:^|&)([^&=]*)=?([^&]*)/g, function ($0, $1, $2) {
    if ($1) {
      data[$1] = $2;
    }
  });
  return data;
}

var Socket$1 = /*#__PURE__*/function (_Emitter) {
  /**
   * Socket constructor.
   *
   * @param {String|Object} uri - uri or options
   * @param {Object} opts - options
   */
  function Socket(uri) {
    var _this;
    var opts = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : {};
    _classCallCheck(this, Socket);
    _this = _callSuper(this, Socket);
    _this.binaryType = defaultBinaryType;
    _this.writeBuffer = [];
    if (uri && "object" === _typeof(uri)) {
      opts = uri;
      uri = null;
    }
    if (uri) {
      uri = parse(uri);
      opts.hostname = uri.host;
      opts.secure = uri.protocol === "https" || uri.protocol === "wss";
      opts.port = uri.port;
      if (uri.query) opts.query = uri.query;
    } else if (opts.host) {
      opts.hostname = parse(opts.host).host;
    }
    installTimerFunctions(_this, opts);
    _this.secure = null != opts.secure ? opts.secure : typeof location !== "undefined" && "https:" === location.protocol;
    if (opts.hostname && !opts.port) {
      // if no port is specified manually, use the protocol default
      opts.port = _this.secure ? "443" : "80";
    }
    _this.hostname = opts.hostname || (typeof location !== "undefined" ? location.hostname : "localhost");
    _this.port = opts.port || (typeof location !== "undefined" && location.port ? location.port : _this.secure ? "443" : "80");
    _this.transports = opts.transports || ["polling", "websocket", "webtransport"];
    _this.writeBuffer = [];
    _this.prevBufferLen = 0;
    _this.opts = Object.assign({
      path: "/engine.io",
      agent: false,
      withCredentials: false,
      upgrade: true,
      timestampParam: "t",
      rememberUpgrade: false,
      addTrailingSlash: true,
      rejectUnauthorized: true,
      perMessageDeflate: {
        threshold: 1024
      },
      transportOptions: {},
      closeOnBeforeunload: false
    }, opts);
    _this.opts.path = _this.opts.path.replace(/\/$/, "") + (_this.opts.addTrailingSlash ? "/" : "");
    if (typeof _this.opts.query === "string") {
      _this.opts.query = decode(_this.opts.query);
    }
    // set on handshake
    _this.id = null;
    _this.upgrades = null;
    _this.pingInterval = null;
    _this.pingTimeout = null;
    // set on heartbeat
    _this.pingTimeoutTimer = null;
    if (typeof addEventListener === "function") {
      if (_this.opts.closeOnBeforeunload) {
        // Firefox closes the connection when the "beforeunload" event is emitted but not Chrome. This event listener
        // ensures every browser behaves the same (no "disconnect" event at the Socket.IO level when the page is
        // closed/reloaded)
        _this.beforeunloadEventListener = function () {
          if (_this.transport) {
            // silently close the transport
            _this.transport.removeAllListeners();
            _this.transport.close();
          }
        };
        addEventListener("beforeunload", _this.beforeunloadEventListener, false);
      }
      if (_this.hostname !== "localhost") {
        _this.offlineEventListener = function () {
          _this.onClose("transport close", {
            description: "network connection lost"
          });
        };
        addEventListener("offline", _this.offlineEventListener, false);
      }
    }
    _this.open();
    return _this;
  }
  /**
   * Creates transport of the given type.
   *
   * @param {String} name - transport name
   * @return {Transport}
   * @private
   */
  _inherits(Socket, _Emitter);
  return _createClass(Socket, [{
    key: "createTransport",
    value: function createTransport(name) {
      var query = Object.assign({}, this.opts.query);
      // append engine.io protocol identifier
      query.EIO = protocol$1;
      // transport name
      query.transport = name;
      // session id if we already have one
      if (this.id) query.sid = this.id;
      var opts = Object.assign({}, this.opts, {
        query: query,
        socket: this,
        hostname: this.hostname,
        secure: this.secure,
        port: this.port
      }, this.opts.transportOptions[name]);
      return new transports[name](opts);
    }
    /**
     * Initializes transport to use and starts probe.
     *
     * @private
     */
  }, {
    key: "open",
    value: function open() {
      var _this2 = this;
      var transport;
      if (this.opts.rememberUpgrade && Socket.priorWebsocketSuccess && this.transports.indexOf("websocket") !== -1) {
        transport = "websocket";
      } else if (0 === this.transports.length) {
        // Emit error on next tick so it can be listened to
        this.setTimeoutFn(function () {
          _this2.emitReserved("error", "No transports available");
        }, 0);
        return;
      } else {
        transport = this.transports[0];
      }
      this.readyState = "opening";
      // Retry with the next transport if the transport is disabled (jsonp: false)
      try {
        transport = this.createTransport(transport);
      } catch (e) {
        this.transports.shift();
        this.open();
        return;
      }
      transport.open();
      this.setTransport(transport);
    }
    /**
     * Sets the current transport. Disables the existing one (if any).
     *
     * @private
     */
  }, {
    key: "setTransport",
    value: function setTransport(transport) {
      var _this3 = this;
      if (this.transport) {
        this.transport.removeAllListeners();
      }
      // set up transport
      this.transport = transport;
      // set up transport listeners
      transport.on("drain", this.onDrain.bind(this)).on("packet", this.onPacket.bind(this)).on("error", this.onError.bind(this)).on("close", function (reason) {
        return _this3.onClose("transport close", reason);
      });
    }
    /**
     * Probes a transport.
     *
     * @param {String} name - transport name
     * @private
     */
  }, {
    key: "probe",
    value: function probe(name) {
      var _this4 = this;
      var transport = this.createTransport(name);
      var failed = false;
      Socket.priorWebsocketSuccess = false;
      var onTransportOpen = function onTransportOpen() {
        if (failed) return;
        transport.send([{
          type: "ping",
          data: "probe"
        }]);
        transport.once("packet", function (msg) {
          if (failed) return;
          if ("pong" === msg.type && "probe" === msg.data) {
            _this4.upgrading = true;
            _this4.emitReserved("upgrading", transport);
            if (!transport) return;
            Socket.priorWebsocketSuccess = "websocket" === transport.name;
            _this4.transport.pause(function () {
              if (failed) return;
              if ("closed" === _this4.readyState) return;
              cleanup();
              _this4.setTransport(transport);
              transport.send([{
                type: "upgrade"
              }]);
              _this4.emitReserved("upgrade", transport);
              transport = null;
              _this4.upgrading = false;
              _this4.flush();
            });
          } else {
            var err = new Error("probe error");
            // @ts-ignore
            err.transport = transport.name;
            _this4.emitReserved("upgradeError", err);
          }
        });
      };
      function freezeTransport() {
        if (failed) return;
        // Any callback called by transport should be ignored since now
        failed = true;
        cleanup();
        transport.close();
        transport = null;
      }
      // Handle any error that happens while probing
      var onerror = function onerror(err) {
        var error = new Error("probe error: " + err);
        // @ts-ignore
        error.transport = transport.name;
        freezeTransport();
        _this4.emitReserved("upgradeError", error);
      };
      function onTransportClose() {
        onerror("transport closed");
      }
      // When the socket is closed while we're probing
      function onclose() {
        onerror("socket closed");
      }
      // When the socket is upgraded while we're probing
      function onupgrade(to) {
        if (transport && to.name !== transport.name) {
          freezeTransport();
        }
      }
      // Remove all listeners on the transport and on self
      var cleanup = function cleanup() {
        transport.removeListener("open", onTransportOpen);
        transport.removeListener("error", onerror);
        transport.removeListener("close", onTransportClose);
        _this4.off("close", onclose);
        _this4.off("upgrading", onupgrade);
      };
      transport.once("open", onTransportOpen);
      transport.once("error", onerror);
      transport.once("close", onTransportClose);
      this.once("close", onclose);
      this.once("upgrading", onupgrade);
      if (this.upgrades.indexOf("webtransport") !== -1 && name !== "webtransport") {
        // favor WebTransport
        this.setTimeoutFn(function () {
          if (!failed) {
            transport.open();
          }
        }, 200);
      } else {
        transport.open();
      }
    }
    /**
     * Called when connection is deemed open.
     *
     * @private
     */
  }, {
    key: "onOpen",
    value: function onOpen() {
      this.readyState = "open";
      Socket.priorWebsocketSuccess = "websocket" === this.transport.name;
      this.emitReserved("open");
      this.flush();
      // we check for `readyState` in case an `open`
      // listener already closed the socket
      if ("open" === this.readyState && this.opts.upgrade) {
        var i = 0;
        var l = this.upgrades.length;
        for (; i < l; i++) {
          this.probe(this.upgrades[i]);
        }
      }
    }
    /**
     * Handles a packet.
     *
     * @private
     */
  }, {
    key: "onPacket",
    value: function onPacket(packet) {
      if ("opening" === this.readyState || "open" === this.readyState || "closing" === this.readyState) {
        this.emitReserved("packet", packet);
        // Socket is live - any packet counts
        this.emitReserved("heartbeat");
        this.resetPingTimeout();
        switch (packet.type) {
          case "open":
            this.onHandshake(JSON.parse(packet.data));
            break;
          case "ping":
            this.sendPacket("pong");
            this.emitReserved("ping");
            this.emitReserved("pong");
            break;
          case "error":
            var err = new Error("server error");
            // @ts-ignore
            err.code = packet.data;
            this.onError(err);
            break;
          case "message":
            this.emitReserved("data", packet.data);
            this.emitReserved("message", packet.data);
            break;
        }
      }
    }
    /**
     * Called upon handshake completion.
     *
     * @param {Object} data - handshake obj
     * @private
     */
  }, {
    key: "onHandshake",
    value: function onHandshake(data) {
      this.emitReserved("handshake", data);
      this.id = data.sid;
      this.transport.query.sid = data.sid;
      this.upgrades = this.filterUpgrades(data.upgrades);
      this.pingInterval = data.pingInterval;
      this.pingTimeout = data.pingTimeout;
      this.maxPayload = data.maxPayload;
      this.onOpen();
      // In case open handler closes socket
      if ("closed" === this.readyState) return;
      this.resetPingTimeout();
    }
    /**
     * Sets and resets ping timeout timer based on server pings.
     *
     * @private
     */
  }, {
    key: "resetPingTimeout",
    value: function resetPingTimeout() {
      var _this5 = this;
      this.clearTimeoutFn(this.pingTimeoutTimer);
      this.pingTimeoutTimer = this.setTimeoutFn(function () {
        _this5.onClose("ping timeout");
      }, this.pingInterval + this.pingTimeout);
      if (this.opts.autoUnref) {
        this.pingTimeoutTimer.unref();
      }
    }
    /**
     * Called on `drain` event
     *
     * @private
     */
  }, {
    key: "onDrain",
    value: function onDrain() {
      this.writeBuffer.splice(0, this.prevBufferLen);
      // setting prevBufferLen = 0 is very important
      // for example, when upgrading, upgrade packet is sent over,
      // and a nonzero prevBufferLen could cause problems on `drain`
      this.prevBufferLen = 0;
      if (0 === this.writeBuffer.length) {
        this.emitReserved("drain");
      } else {
        this.flush();
      }
    }
    /**
     * Flush write buffers.
     *
     * @private
     */
  }, {
    key: "flush",
    value: function flush() {
      if ("closed" !== this.readyState && this.transport.writable && !this.upgrading && this.writeBuffer.length) {
        var packets = this.getWritablePackets();
        this.transport.send(packets);
        // keep track of current length of writeBuffer
        // splice writeBuffer and callbackBuffer on `drain`
        this.prevBufferLen = packets.length;
        this.emitReserved("flush");
      }
    }
    /**
     * Ensure the encoded size of the writeBuffer is below the maxPayload value sent by the server (only for HTTP
     * long-polling)
     *
     * @private
     */
  }, {
    key: "getWritablePackets",
    value: function getWritablePackets() {
      var shouldCheckPayloadSize = this.maxPayload && this.transport.name === "polling" && this.writeBuffer.length > 1;
      if (!shouldCheckPayloadSize) {
        return this.writeBuffer;
      }
      var payloadSize = 1; // first packet type
      for (var i = 0; i < this.writeBuffer.length; i++) {
        var data = this.writeBuffer[i].data;
        if (data) {
          payloadSize += byteLength(data);
        }
        if (i > 0 && payloadSize > this.maxPayload) {
          return this.writeBuffer.slice(0, i);
        }
        payloadSize += 2; // separator + packet type
      }
      return this.writeBuffer;
    }
    /**
     * Sends a message.
     *
     * @param {String} msg - message.
     * @param {Object} options.
     * @param {Function} callback function.
     * @return {Socket} for chaining.
     */
  }, {
    key: "write",
    value: function write(msg, options, fn) {
      this.sendPacket("message", msg, options, fn);
      return this;
    }
  }, {
    key: "send",
    value: function send(msg, options, fn) {
      this.sendPacket("message", msg, options, fn);
      return this;
    }
    /**
     * Sends a packet.
     *
     * @param {String} type: packet type.
     * @param {String} data.
     * @param {Object} options.
     * @param {Function} fn - callback function.
     * @private
     */
  }, {
    key: "sendPacket",
    value: function sendPacket(type, data, options, fn) {
      if ("function" === typeof data) {
        fn = data;
        data = undefined;
      }
      if ("function" === typeof options) {
        fn = options;
        options = null;
      }
      if ("closing" === this.readyState || "closed" === this.readyState) {
        return;
      }
      options = options || {};
      options.compress = false !== options.compress;
      var packet = {
        type: type,
        data: data,
        options: options
      };
      this.emitReserved("packetCreate", packet);
      this.writeBuffer.push(packet);
      if (fn) this.once("flush", fn);
      this.flush();
    }
    /**
     * Closes the connection.
     */
  }, {
    key: "close",
    value: function close() {
      var _this6 = this;
      var close = function close() {
        _this6.onClose("forced close");
        _this6.transport.close();
      };
      var _cleanupAndClose = function cleanupAndClose() {
        _this6.off("upgrade", _cleanupAndClose);
        _this6.off("upgradeError", _cleanupAndClose);
        close();
      };
      var waitForUpgrade = function waitForUpgrade() {
        // wait for upgrade to finish since we can't send packets while pausing a transport
        _this6.once("upgrade", _cleanupAndClose);
        _this6.once("upgradeError", _cleanupAndClose);
      };
      if ("opening" === this.readyState || "open" === this.readyState) {
        this.readyState = "closing";
        if (this.writeBuffer.length) {
          this.once("drain", function () {
            if (_this6.upgrading) {
              waitForUpgrade();
            } else {
              close();
            }
          });
        } else if (this.upgrading) {
          waitForUpgrade();
        } else {
          close();
        }
      }
      return this;
    }
    /**
     * Called upon transport error
     *
     * @private
     */
  }, {
    key: "onError",
    value: function onError(err) {
      Socket.priorWebsocketSuccess = false;
      this.emitReserved("error", err);
      this.onClose("transport error", err);
    }
    /**
     * Called upon transport close.
     *
     * @private
     */
  }, {
    key: "onClose",
    value: function onClose(reason, description) {
      if ("opening" === this.readyState || "open" === this.readyState || "closing" === this.readyState) {
        // clear timers
        this.clearTimeoutFn(this.pingTimeoutTimer);
        // stop event from firing again for transport
        this.transport.removeAllListeners("close");
        // ensure transport won't stay open
        this.transport.close();
        // ignore further transport communication
        this.transport.removeAllListeners();
        if (typeof removeEventListener === "function") {
          removeEventListener("beforeunload", this.beforeunloadEventListener, false);
          removeEventListener("offline", this.offlineEventListener, false);
        }
        // set ready state
        this.readyState = "closed";
        // clear session id
        this.id = null;
        // emit close event
        this.emitReserved("close", reason, description);
        // clean buffers after, so users can still
        // grab the buffers on `close` event
        this.writeBuffer = [];
        this.prevBufferLen = 0;
      }
    }
    /**
     * Filters upgrades, returning only those matching client transports.
     *
     * @param {Array} upgrades - server upgrades
     * @private
     */
  }, {
    key: "filterUpgrades",
    value: function filterUpgrades(upgrades) {
      var filteredUpgrades = [];
      var i = 0;
      var j = upgrades.length;
      for (; i < j; i++) {
        if (~this.transports.indexOf(upgrades[i])) filteredUpgrades.push(upgrades[i]);
      }
      return filteredUpgrades;
    }
  }]);
}(Emitter);
Socket$1.protocol = protocol$1;

Socket$1.protocol;

/**
 * URL parser.
 *
 * @param uri - url
 * @param path - the request path of the connection
 * @param loc - An object meant to mimic window.location.
 *        Defaults to window.location.
 * @public
 */
function url(uri) {
  var path = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : "";
  var loc = arguments.length > 2 ? arguments[2] : undefined;
  var obj = uri;
  // default to window.location
  loc = loc || typeof location !== "undefined" && location;
  if (null == uri) uri = loc.protocol + "//" + loc.host;
  // relative path support
  if (typeof uri === "string") {
    if ("/" === uri.charAt(0)) {
      if ("/" === uri.charAt(1)) {
        uri = loc.protocol + uri;
      } else {
        uri = loc.host + uri;
      }
    }
    if (!/^(https?|wss?):\/\//.test(uri)) {
      if ("undefined" !== typeof loc) {
        uri = loc.protocol + "//" + uri;
      } else {
        uri = "https://" + uri;
      }
    }
    // parse
    obj = parse(uri);
  }
  // make sure we treat `localhost:80` and `localhost` equally
  if (!obj.port) {
    if (/^(http|ws)$/.test(obj.protocol)) {
      obj.port = "80";
    } else if (/^(http|ws)s$/.test(obj.protocol)) {
      obj.port = "443";
    }
  }
  obj.path = obj.path || "/";
  var ipv6 = obj.host.indexOf(":") !== -1;
  var host = ipv6 ? "[" + obj.host + "]" : obj.host;
  // define unique id
  obj.id = obj.protocol + "://" + host + ":" + obj.port + path;
  // define href
  obj.href = obj.protocol + "://" + host + (loc && loc.port === obj.port ? "" : ":" + obj.port);
  return obj;
}

var withNativeArrayBuffer = typeof ArrayBuffer === "function";
var isView = function isView(obj) {
  return typeof ArrayBuffer.isView === "function" ? ArrayBuffer.isView(obj) : obj.buffer instanceof ArrayBuffer;
};
var toString = Object.prototype.toString;
var withNativeBlob = typeof Blob === "function" || typeof Blob !== "undefined" && toString.call(Blob) === "[object BlobConstructor]";
var withNativeFile = typeof File === "function" || typeof File !== "undefined" && toString.call(File) === "[object FileConstructor]";
/**
 * Returns true if obj is a Buffer, an ArrayBuffer, a Blob or a File.
 *
 * @private
 */
function isBinary(obj) {
  return withNativeArrayBuffer && (obj instanceof ArrayBuffer || isView(obj)) || withNativeBlob && obj instanceof Blob || withNativeFile && obj instanceof File;
}
function hasBinary(obj, toJSON) {
  if (!obj || _typeof(obj) !== "object") {
    return false;
  }
  if (Array.isArray(obj)) {
    for (var i = 0, l = obj.length; i < l; i++) {
      if (hasBinary(obj[i])) {
        return true;
      }
    }
    return false;
  }
  if (isBinary(obj)) {
    return true;
  }
  if (obj.toJSON && typeof obj.toJSON === "function" && arguments.length === 1) {
    return hasBinary(obj.toJSON(), true);
  }
  for (var key in obj) {
    if (Object.prototype.hasOwnProperty.call(obj, key) && hasBinary(obj[key])) {
      return true;
    }
  }
  return false;
}

/**
 * Replaces every Buffer | ArrayBuffer | Blob | File in packet with a numbered placeholder.
 *
 * @param {Object} packet - socket.io event packet
 * @return {Object} with deconstructed packet and list of buffers
 * @public
 */
function deconstructPacket(packet) {
  var buffers = [];
  var packetData = packet.data;
  var pack = packet;
  pack.data = _deconstructPacket(packetData, buffers);
  pack.attachments = buffers.length; // number of binary 'attachments'
  return {
    packet: pack,
    buffers: buffers
  };
}
function _deconstructPacket(data, buffers) {
  if (!data) return data;
  if (isBinary(data)) {
    var placeholder = {
      _placeholder: true,
      num: buffers.length
    };
    buffers.push(data);
    return placeholder;
  } else if (Array.isArray(data)) {
    var newData = new Array(data.length);
    for (var i = 0; i < data.length; i++) {
      newData[i] = _deconstructPacket(data[i], buffers);
    }
    return newData;
  } else if (_typeof(data) === "object" && !(data instanceof Date)) {
    var _newData = {};
    for (var key in data) {
      if (Object.prototype.hasOwnProperty.call(data, key)) {
        _newData[key] = _deconstructPacket(data[key], buffers);
      }
    }
    return _newData;
  }
  return data;
}
/**
 * Reconstructs a binary packet from its placeholder packet and buffers
 *
 * @param {Object} packet - event packet with placeholders
 * @param {Array} buffers - binary buffers to put in placeholder positions
 * @return {Object} reconstructed packet
 * @public
 */
function reconstructPacket(packet, buffers) {
  packet.data = _reconstructPacket(packet.data, buffers);
  delete packet.attachments; // no longer useful
  return packet;
}
function _reconstructPacket(data, buffers) {
  if (!data) return data;
  if (data && data._placeholder === true) {
    var isIndexValid = typeof data.num === "number" && data.num >= 0 && data.num < buffers.length;
    if (isIndexValid) {
      return buffers[data.num]; // appropriate buffer (should be natural order anyway)
    } else {
      throw new Error("illegal attachments");
    }
  } else if (Array.isArray(data)) {
    for (var i = 0; i < data.length; i++) {
      data[i] = _reconstructPacket(data[i], buffers);
    }
  } else if (_typeof(data) === "object") {
    for (var key in data) {
      if (Object.prototype.hasOwnProperty.call(data, key)) {
        data[key] = _reconstructPacket(data[key], buffers);
      }
    }
  }
  return data;
}

/**
 * These strings must not be used as event names, as they have a special meaning.
 */
var RESERVED_EVENTS$1 = ["connect", "connect_error", "disconnect", "disconnecting", "newListener", "removeListener" // used by the Node.js EventEmitter
];
/**
 * Protocol version.
 *
 * @public
 */
var protocol = 5;
var PacketType;
(function (PacketType) {
  PacketType[PacketType["CONNECT"] = 0] = "CONNECT";
  PacketType[PacketType["DISCONNECT"] = 1] = "DISCONNECT";
  PacketType[PacketType["EVENT"] = 2] = "EVENT";
  PacketType[PacketType["ACK"] = 3] = "ACK";
  PacketType[PacketType["CONNECT_ERROR"] = 4] = "CONNECT_ERROR";
  PacketType[PacketType["BINARY_EVENT"] = 5] = "BINARY_EVENT";
  PacketType[PacketType["BINARY_ACK"] = 6] = "BINARY_ACK";
})(PacketType || (PacketType = {}));
/**
 * A socket.io Encoder instance
 */
var Encoder = /*#__PURE__*/function () {
  /**
   * Encoder constructor
   *
   * @param {function} replacer - custom replacer to pass down to JSON.parse
   */
  function Encoder(replacer) {
    _classCallCheck(this, Encoder);
    this.replacer = replacer;
  }
  /**
   * Encode a packet as a single string if non-binary, or as a
   * buffer sequence, depending on packet type.
   *
   * @param {Object} obj - packet object
   */
  return _createClass(Encoder, [{
    key: "encode",
    value: function encode(obj) {
      if (obj.type === PacketType.EVENT || obj.type === PacketType.ACK) {
        if (hasBinary(obj)) {
          return this.encodeAsBinary({
            type: obj.type === PacketType.EVENT ? PacketType.BINARY_EVENT : PacketType.BINARY_ACK,
            nsp: obj.nsp,
            data: obj.data,
            id: obj.id
          });
        }
      }
      return [this.encodeAsString(obj)];
    }
    /**
     * Encode packet as string.
     */
  }, {
    key: "encodeAsString",
    value: function encodeAsString(obj) {
      // first is type
      var str = "" + obj.type;
      // attachments if we have them
      if (obj.type === PacketType.BINARY_EVENT || obj.type === PacketType.BINARY_ACK) {
        str += obj.attachments + "-";
      }
      // if we have a namespace other than `/`
      // we append it followed by a comma `,`
      if (obj.nsp && "/" !== obj.nsp) {
        str += obj.nsp + ",";
      }
      // immediately followed by the id
      if (null != obj.id) {
        str += obj.id;
      }
      // json data
      if (null != obj.data) {
        str += JSON.stringify(obj.data, this.replacer);
      }
      return str;
    }
    /**
     * Encode packet as 'buffer sequence' by removing blobs, and
     * deconstructing packet into object with placeholders and
     * a list of buffers.
     */
  }, {
    key: "encodeAsBinary",
    value: function encodeAsBinary(obj) {
      var deconstruction = deconstructPacket(obj);
      var pack = this.encodeAsString(deconstruction.packet);
      var buffers = deconstruction.buffers;
      buffers.unshift(pack); // add packet info to beginning of data list
      return buffers; // write all the buffers
    }
  }]);
}();
// see https://stackoverflow.com/questions/8511281/check-if-a-value-is-an-object-in-javascript
function isObject(value) {
  return Object.prototype.toString.call(value) === "[object Object]";
}
/**
 * A socket.io Decoder instance
 *
 * @return {Object} decoder
 */
var Decoder = /*#__PURE__*/function (_Emitter) {
  /**
   * Decoder constructor
   *
   * @param {function} reviver - custom reviver to pass down to JSON.stringify
   */
  function Decoder(reviver) {
    var _this;
    _classCallCheck(this, Decoder);
    _this = _callSuper(this, Decoder);
    _this.reviver = reviver;
    return _this;
  }
  /**
   * Decodes an encoded packet string into packet JSON.
   *
   * @param {String} obj - encoded packet
   */
  _inherits(Decoder, _Emitter);
  return _createClass(Decoder, [{
    key: "add",
    value: function add(obj) {
      var packet;
      if (typeof obj === "string") {
        if (this.reconstructor) {
          throw new Error("got plaintext data when reconstructing a packet");
        }
        packet = this.decodeString(obj);
        var isBinaryEvent = packet.type === PacketType.BINARY_EVENT;
        if (isBinaryEvent || packet.type === PacketType.BINARY_ACK) {
          packet.type = isBinaryEvent ? PacketType.EVENT : PacketType.ACK;
          // binary packet's json
          this.reconstructor = new BinaryReconstructor(packet);
          // no attachments, labeled binary but no binary data to follow
          if (packet.attachments === 0) {
            _superPropGet(Decoder, "emitReserved", this, 3)(["decoded", packet]);
          }
        } else {
          // non-binary full packet
          _superPropGet(Decoder, "emitReserved", this, 3)(["decoded", packet]);
        }
      } else if (isBinary(obj) || obj.base64) {
        // raw binary data
        if (!this.reconstructor) {
          throw new Error("got binary data when not reconstructing a packet");
        } else {
          packet = this.reconstructor.takeBinaryData(obj);
          if (packet) {
            // received final buffer
            this.reconstructor = null;
            _superPropGet(Decoder, "emitReserved", this, 3)(["decoded", packet]);
          }
        }
      } else {
        throw new Error("Unknown type: " + obj);
      }
    }
    /**
     * Decode a packet String (JSON data)
     *
     * @param {String} str
     * @return {Object} packet
     */
  }, {
    key: "decodeString",
    value: function decodeString(str) {
      var i = 0;
      // look up type
      var p = {
        type: Number(str.charAt(0))
      };
      if (PacketType[p.type] === undefined) {
        throw new Error("unknown packet type " + p.type);
      }
      // look up attachments if type binary
      if (p.type === PacketType.BINARY_EVENT || p.type === PacketType.BINARY_ACK) {
        var start = i + 1;
        while (str.charAt(++i) !== "-" && i != str.length) {}
        var buf = str.substring(start, i);
        if (buf != Number(buf) || str.charAt(i) !== "-") {
          throw new Error("Illegal attachments");
        }
        p.attachments = Number(buf);
      }
      // look up namespace (if any)
      if ("/" === str.charAt(i + 1)) {
        var _start = i + 1;
        while (++i) {
          var c = str.charAt(i);
          if ("," === c) break;
          if (i === str.length) break;
        }
        p.nsp = str.substring(_start, i);
      } else {
        p.nsp = "/";
      }
      // look up id
      var next = str.charAt(i + 1);
      if ("" !== next && Number(next) == next) {
        var _start2 = i + 1;
        while (++i) {
          var _c = str.charAt(i);
          if (null == _c || Number(_c) != _c) {
            --i;
            break;
          }
          if (i === str.length) break;
        }
        p.id = Number(str.substring(_start2, i + 1));
      }
      // look up json data
      if (str.charAt(++i)) {
        var payload = this.tryParse(str.substr(i));
        if (Decoder.isPayloadValid(p.type, payload)) {
          p.data = payload;
        } else {
          throw new Error("invalid payload");
        }
      }
      return p;
    }
  }, {
    key: "tryParse",
    value: function tryParse(str) {
      try {
        return JSON.parse(str, this.reviver);
      } catch (e) {
        return false;
      }
    }
  }, {
    key: "destroy",
    value:
    /**
     * Deallocates a parser's resources
     */
    function destroy() {
      if (this.reconstructor) {
        this.reconstructor.finishedReconstruction();
        this.reconstructor = null;
      }
    }
  }], [{
    key: "isPayloadValid",
    value: function isPayloadValid(type, payload) {
      switch (type) {
        case PacketType.CONNECT:
          return isObject(payload);
        case PacketType.DISCONNECT:
          return payload === undefined;
        case PacketType.CONNECT_ERROR:
          return typeof payload === "string" || isObject(payload);
        case PacketType.EVENT:
        case PacketType.BINARY_EVENT:
          return Array.isArray(payload) && (typeof payload[0] === "number" || typeof payload[0] === "string" && RESERVED_EVENTS$1.indexOf(payload[0]) === -1);
        case PacketType.ACK:
        case PacketType.BINARY_ACK:
          return Array.isArray(payload);
      }
    }
  }]);
}(Emitter);
/**
 * A manager of a binary event's 'buffer sequence'. Should
 * be constructed whenever a packet of type BINARY_EVENT is
 * decoded.
 *
 * @param {Object} packet
 * @return {BinaryReconstructor} initialized reconstructor
 */
var BinaryReconstructor = /*#__PURE__*/function () {
  function BinaryReconstructor(packet) {
    _classCallCheck(this, BinaryReconstructor);
    this.packet = packet;
    this.buffers = [];
    this.reconPack = packet;
  }
  /**
   * Method to be called when binary data received from connection
   * after a BINARY_EVENT packet.
   *
   * @param {Buffer | ArrayBuffer} binData - the raw binary data received
   * @return {null | Object} returns null if more binary data is expected or
   *   a reconstructed packet object if all buffers have been received.
   */
  return _createClass(BinaryReconstructor, [{
    key: "takeBinaryData",
    value: function takeBinaryData(binData) {
      this.buffers.push(binData);
      if (this.buffers.length === this.reconPack.attachments) {
        // done with buffer list
        var packet = reconstructPacket(this.reconPack, this.buffers);
        this.finishedReconstruction();
        return packet;
      }
      return null;
    }
    /**
     * Cleans up binary packet reconstruction variables.
     */
  }, {
    key: "finishedReconstruction",
    value: function finishedReconstruction() {
      this.reconPack = null;
      this.buffers = [];
    }
  }]);
}();

var parser = /*#__PURE__*/Object.freeze({
    __proto__: null,
    Decoder: Decoder,
    Encoder: Encoder,
    get PacketType () { return PacketType; },
    protocol: protocol
});

function on(obj, ev, fn) {
  obj.on(ev, fn);
  return function subDestroy() {
    obj.off(ev, fn);
  };
}

/**
 * Internal events.
 * These events can't be emitted by the user.
 */
var RESERVED_EVENTS = Object.freeze({
  connect: 1,
  connect_error: 1,
  disconnect: 1,
  disconnecting: 1,
  // EventEmitter reserved events: https://nodejs.org/api/events.html#events_event_newlistener
  newListener: 1,
  removeListener: 1
});
/**
 * A Socket is the fundamental class for interacting with the server.
 *
 * A Socket belongs to a certain Namespace (by default /) and uses an underlying {@link Manager} to communicate.
 *
 * @example
 * const socket = io();
 *
 * socket.on("connect", () => {
 *   console.log("connected");
 * });
 *
 * // send an event to the server
 * socket.emit("foo", "bar");
 *
 * socket.on("foobar", () => {
 *   // an event was received from the server
 * });
 *
 * // upon disconnection
 * socket.on("disconnect", (reason) => {
 *   console.log(`disconnected due to ${reason}`);
 * });
 */
var Socket = /*#__PURE__*/function (_Emitter) {
  /**
   * `Socket` constructor.
   */
  function Socket(io, nsp, opts) {
    var _this;
    _classCallCheck(this, Socket);
    _this = _callSuper(this, Socket);
    /**
     * Whether the socket is currently connected to the server.
     *
     * @example
     * const socket = io();
     *
     * socket.on("connect", () => {
     *   console.log(socket.connected); // true
     * });
     *
     * socket.on("disconnect", () => {
     *   console.log(socket.connected); // false
     * });
     */
    _this.connected = false;
    /**
     * Whether the connection state was recovered after a temporary disconnection. In that case, any missed packets will
     * be transmitted by the server.
     */
    _this.recovered = false;
    /**
     * Buffer for packets received before the CONNECT packet
     */
    _this.receiveBuffer = [];
    /**
     * Buffer for packets that will be sent once the socket is connected
     */
    _this.sendBuffer = [];
    /**
     * The queue of packets to be sent with retry in case of failure.
     *
     * Packets are sent one by one, each waiting for the server acknowledgement, in order to guarantee the delivery order.
     * @private
     */
    _this._queue = [];
    /**
     * A sequence to generate the ID of the {@link QueuedPacket}.
     * @private
     */
    _this._queueSeq = 0;
    _this.ids = 0;
    /**
     * A map containing acknowledgement handlers.
     *
     * The `withError` attribute is used to differentiate handlers that accept an error as first argument:
     *
     * - `socket.emit("test", (err, value) => { ... })` with `ackTimeout` option
     * - `socket.timeout(5000).emit("test", (err, value) => { ... })`
     * - `const value = await socket.emitWithAck("test")`
     *
     * From those that don't:
     *
     * - `socket.emit("test", (value) => { ... });`
     *
     * In the first case, the handlers will be called with an error when:
     *
     * - the timeout is reached
     * - the socket gets disconnected
     *
     * In the second case, the handlers will be simply discarded upon disconnection, since the client will never receive
     * an acknowledgement from the server.
     *
     * @private
     */
    _this.acks = {};
    _this.flags = {};
    _this.io = io;
    _this.nsp = nsp;
    if (opts && opts.auth) {
      _this.auth = opts.auth;
    }
    _this._opts = Object.assign({}, opts);
    if (_this.io._autoConnect) _this.open();
    return _this;
  }
  /**
   * Whether the socket is currently disconnected
   *
   * @example
   * const socket = io();
   *
   * socket.on("connect", () => {
   *   console.log(socket.disconnected); // false
   * });
   *
   * socket.on("disconnect", () => {
   *   console.log(socket.disconnected); // true
   * });
   */
  _inherits(Socket, _Emitter);
  return _createClass(Socket, [{
    key: "disconnected",
    get: function get() {
      return !this.connected;
    }
    /**
     * Subscribe to open, close and packet events
     *
     * @private
     */
  }, {
    key: "subEvents",
    value: function subEvents() {
      if (this.subs) return;
      var io = this.io;
      this.subs = [on(io, "open", this.onopen.bind(this)), on(io, "packet", this.onpacket.bind(this)), on(io, "error", this.onerror.bind(this)), on(io, "close", this.onclose.bind(this))];
    }
    /**
     * Whether the Socket will try to reconnect when its Manager connects or reconnects.
     *
     * @example
     * const socket = io();
     *
     * console.log(socket.active); // true
     *
     * socket.on("disconnect", (reason) => {
     *   if (reason === "io server disconnect") {
     *     // the disconnection was initiated by the server, you need to manually reconnect
     *     console.log(socket.active); // false
     *   }
     *   // else the socket will automatically try to reconnect
     *   console.log(socket.active); // true
     * });
     */
  }, {
    key: "active",
    get: function get() {
      return !!this.subs;
    }
    /**
     * "Opens" the socket.
     *
     * @example
     * const socket = io({
     *   autoConnect: false
     * });
     *
     * socket.connect();
     */
  }, {
    key: "connect",
    value: function connect() {
      if (this.connected) return this;
      this.subEvents();
      if (!this.io["_reconnecting"]) this.io.open(); // ensure open
      if ("open" === this.io._readyState) this.onopen();
      return this;
    }
    /**
     * Alias for {@link connect()}.
     */
  }, {
    key: "open",
    value: function open() {
      return this.connect();
    }
    /**
     * Sends a `message` event.
     *
     * This method mimics the WebSocket.send() method.
     *
     * @see https://developer.mozilla.org/en-US/docs/Web/API/WebSocket/send
     *
     * @example
     * socket.send("hello");
     *
     * // this is equivalent to
     * socket.emit("message", "hello");
     *
     * @return self
     */
  }, {
    key: "send",
    value: function send() {
      for (var _len = arguments.length, args = new Array(_len), _key = 0; _key < _len; _key++) {
        args[_key] = arguments[_key];
      }
      args.unshift("message");
      this.emit.apply(this, args);
      return this;
    }
    /**
     * Override `emit`.
     * If the event is in `events`, it's emitted normally.
     *
     * @example
     * socket.emit("hello", "world");
     *
     * // all serializable datastructures are supported (no need to call JSON.stringify)
     * socket.emit("hello", 1, "2", { 3: ["4"], 5: Uint8Array.from([6]) });
     *
     * // with an acknowledgement from the server
     * socket.emit("hello", "world", (val) => {
     *   // ...
     * });
     *
     * @return self
     */
  }, {
    key: "emit",
    value: function emit(ev) {
      if (RESERVED_EVENTS.hasOwnProperty(ev)) {
        throw new Error('"' + ev.toString() + '" is a reserved event name');
      }
      for (var _len2 = arguments.length, args = new Array(_len2 > 1 ? _len2 - 1 : 0), _key2 = 1; _key2 < _len2; _key2++) {
        args[_key2 - 1] = arguments[_key2];
      }
      args.unshift(ev);
      if (this._opts.retries && !this.flags.fromQueue && !this.flags["volatile"]) {
        this._addToQueue(args);
        return this;
      }
      var packet = {
        type: PacketType.EVENT,
        data: args
      };
      packet.options = {};
      packet.options.compress = this.flags.compress !== false;
      // event ack callback
      if ("function" === typeof args[args.length - 1]) {
        var id = this.ids++;
        var ack = args.pop();
        this._registerAckCallback(id, ack);
        packet.id = id;
      }
      var isTransportWritable = this.io.engine && this.io.engine.transport && this.io.engine.transport.writable;
      var discardPacket = this.flags["volatile"] && (!isTransportWritable || !this.connected);
      if (discardPacket) ; else if (this.connected) {
        this.notifyOutgoingListeners(packet);
        this.packet(packet);
      } else {
        this.sendBuffer.push(packet);
      }
      this.flags = {};
      return this;
    }
    /**
     * @private
     */
  }, {
    key: "_registerAckCallback",
    value: function _registerAckCallback(id, ack) {
      var _this2 = this;
      var _a;
      var timeout = (_a = this.flags.timeout) !== null && _a !== void 0 ? _a : this._opts.ackTimeout;
      if (timeout === undefined) {
        this.acks[id] = ack;
        return;
      }
      // @ts-ignore
      var timer = this.io.setTimeoutFn(function () {
        delete _this2.acks[id];
        for (var i = 0; i < _this2.sendBuffer.length; i++) {
          if (_this2.sendBuffer[i].id === id) {
            _this2.sendBuffer.splice(i, 1);
          }
        }
        ack.call(_this2, new Error("operation has timed out"));
      }, timeout);
      var fn = function fn() {
        // @ts-ignore
        _this2.io.clearTimeoutFn(timer);
        for (var _len3 = arguments.length, args = new Array(_len3), _key3 = 0; _key3 < _len3; _key3++) {
          args[_key3] = arguments[_key3];
        }
        ack.apply(_this2, args);
      };
      fn.withError = true;
      this.acks[id] = fn;
    }
    /**
     * Emits an event and waits for an acknowledgement
     *
     * @example
     * // without timeout
     * const response = await socket.emitWithAck("hello", "world");
     *
     * // with a specific timeout
     * try {
     *   const response = await socket.timeout(1000).emitWithAck("hello", "world");
     * } catch (err) {
     *   // the server did not acknowledge the event in the given delay
     * }
     *
     * @return a Promise that will be fulfilled when the server acknowledges the event
     */
  }, {
    key: "emitWithAck",
    value: function emitWithAck(ev) {
      var _this3 = this;
      for (var _len4 = arguments.length, args = new Array(_len4 > 1 ? _len4 - 1 : 0), _key4 = 1; _key4 < _len4; _key4++) {
        args[_key4 - 1] = arguments[_key4];
      }
      return new Promise(function (resolve, reject) {
        var fn = function fn(arg1, arg2) {
          return arg1 ? reject(arg1) : resolve(arg2);
        };
        fn.withError = true;
        args.push(fn);
        _this3.emit.apply(_this3, [ev].concat(args));
      });
    }
    /**
     * Add the packet to the queue.
     * @param args
     * @private
     */
  }, {
    key: "_addToQueue",
    value: function _addToQueue(args) {
      var _this4 = this;
      var ack;
      if (typeof args[args.length - 1] === "function") {
        ack = args.pop();
      }
      var packet = {
        id: this._queueSeq++,
        tryCount: 0,
        pending: false,
        args: args,
        flags: Object.assign({
          fromQueue: true
        }, this.flags)
      };
      args.push(function (err) {
        if (packet !== _this4._queue[0]) {
          // the packet has already been acknowledged
          return;
        }
        var hasError = err !== null;
        if (hasError) {
          if (packet.tryCount > _this4._opts.retries) {
            _this4._queue.shift();
            if (ack) {
              ack(err);
            }
          }
        } else {
          _this4._queue.shift();
          if (ack) {
            for (var _len5 = arguments.length, responseArgs = new Array(_len5 > 1 ? _len5 - 1 : 0), _key5 = 1; _key5 < _len5; _key5++) {
              responseArgs[_key5 - 1] = arguments[_key5];
            }
            ack.apply(void 0, [null].concat(responseArgs));
          }
        }
        packet.pending = false;
        return _this4._drainQueue();
      });
      this._queue.push(packet);
      this._drainQueue();
    }
    /**
     * Send the first packet of the queue, and wait for an acknowledgement from the server.
     * @param force - whether to resend a packet that has not been acknowledged yet
     *
     * @private
     */
  }, {
    key: "_drainQueue",
    value: function _drainQueue() {
      var force = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : false;
      if (!this.connected || this._queue.length === 0) {
        return;
      }
      var packet = this._queue[0];
      if (packet.pending && !force) {
        return;
      }
      packet.pending = true;
      packet.tryCount++;
      this.flags = packet.flags;
      this.emit.apply(this, packet.args);
    }
    /**
     * Sends a packet.
     *
     * @param packet
     * @private
     */
  }, {
    key: "packet",
    value: function packet(_packet) {
      _packet.nsp = this.nsp;
      this.io._packet(_packet);
    }
    /**
     * Called upon engine `open`.
     *
     * @private
     */
  }, {
    key: "onopen",
    value: function onopen() {
      var _this5 = this;
      if (typeof this.auth == "function") {
        this.auth(function (data) {
          _this5._sendConnectPacket(data);
        });
      } else {
        this._sendConnectPacket(this.auth);
      }
    }
    /**
     * Sends a CONNECT packet to initiate the Socket.IO session.
     *
     * @param data
     * @private
     */
  }, {
    key: "_sendConnectPacket",
    value: function _sendConnectPacket(data) {
      this.packet({
        type: PacketType.CONNECT,
        data: this._pid ? Object.assign({
          pid: this._pid,
          offset: this._lastOffset
        }, data) : data
      });
    }
    /**
     * Called upon engine or manager `error`.
     *
     * @param err
     * @private
     */
  }, {
    key: "onerror",
    value: function onerror(err) {
      if (!this.connected) {
        this.emitReserved("connect_error", err);
      }
    }
    /**
     * Called upon engine `close`.
     *
     * @param reason
     * @param description
     * @private
     */
  }, {
    key: "onclose",
    value: function onclose(reason, description) {
      this.connected = false;
      delete this.id;
      this.emitReserved("disconnect", reason, description);
      this._clearAcks();
    }
    /**
     * Clears the acknowledgement handlers upon disconnection, since the client will never receive an acknowledgement from
     * the server.
     *
     * @private
     */
  }, {
    key: "_clearAcks",
    value: function _clearAcks() {
      var _this6 = this;
      Object.keys(this.acks).forEach(function (id) {
        var isBuffered = _this6.sendBuffer.some(function (packet) {
          return String(packet.id) === id;
        });
        if (!isBuffered) {
          // note: handlers that do not accept an error as first argument are ignored here
          var ack = _this6.acks[id];
          delete _this6.acks[id];
          if (ack.withError) {
            ack.call(_this6, new Error("socket has been disconnected"));
          }
        }
      });
    }
    /**
     * Called with socket packet.
     *
     * @param packet
     * @private
     */
  }, {
    key: "onpacket",
    value: function onpacket(packet) {
      var sameNamespace = packet.nsp === this.nsp;
      if (!sameNamespace) return;
      switch (packet.type) {
        case PacketType.CONNECT:
          if (packet.data && packet.data.sid) {
            this.onconnect(packet.data.sid, packet.data.pid);
          } else {
            this.emitReserved("connect_error", new Error("It seems you are trying to reach a Socket.IO server in v2.x with a v3.x client, but they are not compatible (more information here: https://socket.io/docs/v3/migrating-from-2-x-to-3-0/)"));
          }
          break;
        case PacketType.EVENT:
        case PacketType.BINARY_EVENT:
          this.onevent(packet);
          break;
        case PacketType.ACK:
        case PacketType.BINARY_ACK:
          this.onack(packet);
          break;
        case PacketType.DISCONNECT:
          this.ondisconnect();
          break;
        case PacketType.CONNECT_ERROR:
          this.destroy();
          var err = new Error(packet.data.message);
          // @ts-ignore
          err.data = packet.data.data;
          this.emitReserved("connect_error", err);
          break;
      }
    }
    /**
     * Called upon a server event.
     *
     * @param packet
     * @private
     */
  }, {
    key: "onevent",
    value: function onevent(packet) {
      var args = packet.data || [];
      if (null != packet.id) {
        args.push(this.ack(packet.id));
      }
      if (this.connected) {
        this.emitEvent(args);
      } else {
        this.receiveBuffer.push(Object.freeze(args));
      }
    }
  }, {
    key: "emitEvent",
    value: function emitEvent(args) {
      if (this._anyListeners && this._anyListeners.length) {
        var listeners = this._anyListeners.slice();
        var _iterator = _createForOfIteratorHelper(listeners),
          _step;
        try {
          for (_iterator.s(); !(_step = _iterator.n()).done;) {
            var listener = _step.value;
            listener.apply(this, args);
          }
        } catch (err) {
          _iterator.e(err);
        } finally {
          _iterator.f();
        }
      }
      _superPropGet(Socket, "emit", this, 1).apply(this, args);
      if (this._pid && args.length && typeof args[args.length - 1] === "string") {
        this._lastOffset = args[args.length - 1];
      }
    }
    /**
     * Produces an ack callback to emit with an event.
     *
     * @private
     */
  }, {
    key: "ack",
    value: function ack(id) {
      var self = this;
      var sent = false;
      return function () {
        // prevent double callbacks
        if (sent) return;
        sent = true;
        for (var _len6 = arguments.length, args = new Array(_len6), _key6 = 0; _key6 < _len6; _key6++) {
          args[_key6] = arguments[_key6];
        }
        self.packet({
          type: PacketType.ACK,
          id: id,
          data: args
        });
      };
    }
    /**
     * Called upon a server acknowledgement.
     *
     * @param packet
     * @private
     */
  }, {
    key: "onack",
    value: function onack(packet) {
      var ack = this.acks[packet.id];
      if (typeof ack !== "function") {
        return;
      }
      delete this.acks[packet.id];
      // @ts-ignore FIXME ack is incorrectly inferred as 'never'
      if (ack.withError) {
        packet.data.unshift(null);
      }
      // @ts-ignore
      ack.apply(this, packet.data);
    }
    /**
     * Called upon server connect.
     *
     * @private
     */
  }, {
    key: "onconnect",
    value: function onconnect(id, pid) {
      this.id = id;
      this.recovered = pid && this._pid === pid;
      this._pid = pid; // defined only if connection state recovery is enabled
      this.connected = true;
      this.emitBuffered();
      this.emitReserved("connect");
      this._drainQueue(true);
    }
    /**
     * Emit buffered events (received and emitted).
     *
     * @private
     */
  }, {
    key: "emitBuffered",
    value: function emitBuffered() {
      var _this7 = this;
      this.receiveBuffer.forEach(function (args) {
        return _this7.emitEvent(args);
      });
      this.receiveBuffer = [];
      this.sendBuffer.forEach(function (packet) {
        _this7.notifyOutgoingListeners(packet);
        _this7.packet(packet);
      });
      this.sendBuffer = [];
    }
    /**
     * Called upon server disconnect.
     *
     * @private
     */
  }, {
    key: "ondisconnect",
    value: function ondisconnect() {
      this.destroy();
      this.onclose("io server disconnect");
    }
    /**
     * Called upon forced client/server side disconnections,
     * this method ensures the manager stops tracking us and
     * that reconnections don't get triggered for this.
     *
     * @private
     */
  }, {
    key: "destroy",
    value: function destroy() {
      if (this.subs) {
        // clean subscriptions to avoid reconnections
        this.subs.forEach(function (subDestroy) {
          return subDestroy();
        });
        this.subs = undefined;
      }
      this.io["_destroy"](this);
    }
    /**
     * Disconnects the socket manually. In that case, the socket will not try to reconnect.
     *
     * If this is the last active Socket instance of the {@link Manager}, the low-level connection will be closed.
     *
     * @example
     * const socket = io();
     *
     * socket.on("disconnect", (reason) => {
     *   // console.log(reason); prints "io client disconnect"
     * });
     *
     * socket.disconnect();
     *
     * @return self
     */
  }, {
    key: "disconnect",
    value: function disconnect() {
      if (this.connected) {
        this.packet({
          type: PacketType.DISCONNECT
        });
      }
      // remove socket from pool
      this.destroy();
      if (this.connected) {
        // fire events
        this.onclose("io client disconnect");
      }
      return this;
    }
    /**
     * Alias for {@link disconnect()}.
     *
     * @return self
     */
  }, {
    key: "close",
    value: function close() {
      return this.disconnect();
    }
    /**
     * Sets the compress flag.
     *
     * @example
     * socket.compress(false).emit("hello");
     *
     * @param compress - if `true`, compresses the sending data
     * @return self
     */
  }, {
    key: "compress",
    value: function compress(_compress) {
      this.flags.compress = _compress;
      return this;
    }
    /**
     * Sets a modifier for a subsequent event emission that the event message will be dropped when this socket is not
     * ready to send messages.
     *
     * @example
     * socket.volatile.emit("hello"); // the server may or may not receive it
     *
     * @returns self
     */
  }, {
    key: "volatile",
    get: function get() {
      this.flags["volatile"] = true;
      return this;
    }
    /**
     * Sets a modifier for a subsequent event emission that the callback will be called with an error when the
     * given number of milliseconds have elapsed without an acknowledgement from the server:
     *
     * @example
     * socket.timeout(5000).emit("my-event", (err) => {
     *   if (err) {
     *     // the server did not acknowledge the event in the given delay
     *   }
     * });
     *
     * @returns self
     */
  }, {
    key: "timeout",
    value: function timeout(_timeout) {
      this.flags.timeout = _timeout;
      return this;
    }
    /**
     * Adds a listener that will be fired when any event is emitted. The event name is passed as the first argument to the
     * callback.
     *
     * @example
     * socket.onAny((event, ...args) => {
     *   console.log(`got ${event}`);
     * });
     *
     * @param listener
     */
  }, {
    key: "onAny",
    value: function onAny(listener) {
      this._anyListeners = this._anyListeners || [];
      this._anyListeners.push(listener);
      return this;
    }
    /**
     * Adds a listener that will be fired when any event is emitted. The event name is passed as the first argument to the
     * callback. The listener is added to the beginning of the listeners array.
     *
     * @example
     * socket.prependAny((event, ...args) => {
     *   console.log(`got event ${event}`);
     * });
     *
     * @param listener
     */
  }, {
    key: "prependAny",
    value: function prependAny(listener) {
      this._anyListeners = this._anyListeners || [];
      this._anyListeners.unshift(listener);
      return this;
    }
    /**
     * Removes the listener that will be fired when any event is emitted.
     *
     * @example
     * const catchAllListener = (event, ...args) => {
     *   console.log(`got event ${event}`);
     * }
     *
     * socket.onAny(catchAllListener);
     *
     * // remove a specific listener
     * socket.offAny(catchAllListener);
     *
     * // or remove all listeners
     * socket.offAny();
     *
     * @param listener
     */
  }, {
    key: "offAny",
    value: function offAny(listener) {
      if (!this._anyListeners) {
        return this;
      }
      if (listener) {
        var listeners = this._anyListeners;
        for (var i = 0; i < listeners.length; i++) {
          if (listener === listeners[i]) {
            listeners.splice(i, 1);
            return this;
          }
        }
      } else {
        this._anyListeners = [];
      }
      return this;
    }
    /**
     * Returns an array of listeners that are listening for any event that is specified. This array can be manipulated,
     * e.g. to remove listeners.
     */
  }, {
    key: "listenersAny",
    value: function listenersAny() {
      return this._anyListeners || [];
    }
    /**
     * Adds a listener that will be fired when any event is emitted. The event name is passed as the first argument to the
     * callback.
     *
     * Note: acknowledgements sent to the server are not included.
     *
     * @example
     * socket.onAnyOutgoing((event, ...args) => {
     *   console.log(`sent event ${event}`);
     * });
     *
     * @param listener
     */
  }, {
    key: "onAnyOutgoing",
    value: function onAnyOutgoing(listener) {
      this._anyOutgoingListeners = this._anyOutgoingListeners || [];
      this._anyOutgoingListeners.push(listener);
      return this;
    }
    /**
     * Adds a listener that will be fired when any event is emitted. The event name is passed as the first argument to the
     * callback. The listener is added to the beginning of the listeners array.
     *
     * Note: acknowledgements sent to the server are not included.
     *
     * @example
     * socket.prependAnyOutgoing((event, ...args) => {
     *   console.log(`sent event ${event}`);
     * });
     *
     * @param listener
     */
  }, {
    key: "prependAnyOutgoing",
    value: function prependAnyOutgoing(listener) {
      this._anyOutgoingListeners = this._anyOutgoingListeners || [];
      this._anyOutgoingListeners.unshift(listener);
      return this;
    }
    /**
     * Removes the listener that will be fired when any event is emitted.
     *
     * @example
     * const catchAllListener = (event, ...args) => {
     *   console.log(`sent event ${event}`);
     * }
     *
     * socket.onAnyOutgoing(catchAllListener);
     *
     * // remove a specific listener
     * socket.offAnyOutgoing(catchAllListener);
     *
     * // or remove all listeners
     * socket.offAnyOutgoing();
     *
     * @param [listener] - the catch-all listener (optional)
     */
  }, {
    key: "offAnyOutgoing",
    value: function offAnyOutgoing(listener) {
      if (!this._anyOutgoingListeners) {
        return this;
      }
      if (listener) {
        var listeners = this._anyOutgoingListeners;
        for (var i = 0; i < listeners.length; i++) {
          if (listener === listeners[i]) {
            listeners.splice(i, 1);
            return this;
          }
        }
      } else {
        this._anyOutgoingListeners = [];
      }
      return this;
    }
    /**
     * Returns an array of listeners that are listening for any event that is specified. This array can be manipulated,
     * e.g. to remove listeners.
     */
  }, {
    key: "listenersAnyOutgoing",
    value: function listenersAnyOutgoing() {
      return this._anyOutgoingListeners || [];
    }
    /**
     * Notify the listeners for each packet sent
     *
     * @param packet
     *
     * @private
     */
  }, {
    key: "notifyOutgoingListeners",
    value: function notifyOutgoingListeners(packet) {
      if (this._anyOutgoingListeners && this._anyOutgoingListeners.length) {
        var listeners = this._anyOutgoingListeners.slice();
        var _iterator2 = _createForOfIteratorHelper(listeners),
          _step2;
        try {
          for (_iterator2.s(); !(_step2 = _iterator2.n()).done;) {
            var listener = _step2.value;
            listener.apply(this, packet.data);
          }
        } catch (err) {
          _iterator2.e(err);
        } finally {
          _iterator2.f();
        }
      }
    }
  }]);
}(Emitter);

/**
 * Initialize backoff timer with `opts`.
 *
 * - `min` initial timeout in milliseconds [100]
 * - `max` max timeout [10000]
 * - `jitter` [0]
 * - `factor` [2]
 *
 * @param {Object} opts
 * @api public
 */
function Backoff(opts) {
  opts = opts || {};
  this.ms = opts.min || 100;
  this.max = opts.max || 10000;
  this.factor = opts.factor || 2;
  this.jitter = opts.jitter > 0 && opts.jitter <= 1 ? opts.jitter : 0;
  this.attempts = 0;
}
/**
 * Return the backoff duration.
 *
 * @return {Number}
 * @api public
 */
Backoff.prototype.duration = function () {
  var ms = this.ms * Math.pow(this.factor, this.attempts++);
  if (this.jitter) {
    var rand = Math.random();
    var deviation = Math.floor(rand * this.jitter * ms);
    ms = (Math.floor(rand * 10) & 1) == 0 ? ms - deviation : ms + deviation;
  }
  return Math.min(ms, this.max) | 0;
};
/**
 * Reset the number of attempts.
 *
 * @api public
 */
Backoff.prototype.reset = function () {
  this.attempts = 0;
};
/**
 * Set the minimum duration
 *
 * @api public
 */
Backoff.prototype.setMin = function (min) {
  this.ms = min;
};
/**
 * Set the maximum duration
 *
 * @api public
 */
Backoff.prototype.setMax = function (max) {
  this.max = max;
};
/**
 * Set the jitter
 *
 * @api public
 */
Backoff.prototype.setJitter = function (jitter) {
  this.jitter = jitter;
};

var Manager = /*#__PURE__*/function (_Emitter) {
  function Manager(uri, opts) {
    var _this;
    _classCallCheck(this, Manager);
    var _a;
    _this = _callSuper(this, Manager);
    _this.nsps = {};
    _this.subs = [];
    if (uri && "object" === _typeof(uri)) {
      opts = uri;
      uri = undefined;
    }
    opts = opts || {};
    opts.path = opts.path || "/socket.io";
    _this.opts = opts;
    installTimerFunctions(_this, opts);
    _this.reconnection(opts.reconnection !== false);
    _this.reconnectionAttempts(opts.reconnectionAttempts || Infinity);
    _this.reconnectionDelay(opts.reconnectionDelay || 1000);
    _this.reconnectionDelayMax(opts.reconnectionDelayMax || 5000);
    _this.randomizationFactor((_a = opts.randomizationFactor) !== null && _a !== void 0 ? _a : 0.5);
    _this.backoff = new Backoff({
      min: _this.reconnectionDelay(),
      max: _this.reconnectionDelayMax(),
      jitter: _this.randomizationFactor()
    });
    _this.timeout(null == opts.timeout ? 20000 : opts.timeout);
    _this._readyState = "closed";
    _this.uri = uri;
    var _parser = opts.parser || parser;
    _this.encoder = new _parser.Encoder();
    _this.decoder = new _parser.Decoder();
    _this._autoConnect = opts.autoConnect !== false;
    if (_this._autoConnect) _this.open();
    return _this;
  }
  _inherits(Manager, _Emitter);
  return _createClass(Manager, [{
    key: "reconnection",
    value: function reconnection(v) {
      if (!arguments.length) return this._reconnection;
      this._reconnection = !!v;
      return this;
    }
  }, {
    key: "reconnectionAttempts",
    value: function reconnectionAttempts(v) {
      if (v === undefined) return this._reconnectionAttempts;
      this._reconnectionAttempts = v;
      return this;
    }
  }, {
    key: "reconnectionDelay",
    value: function reconnectionDelay(v) {
      var _a;
      if (v === undefined) return this._reconnectionDelay;
      this._reconnectionDelay = v;
      (_a = this.backoff) === null || _a === void 0 ? void 0 : _a.setMin(v);
      return this;
    }
  }, {
    key: "randomizationFactor",
    value: function randomizationFactor(v) {
      var _a;
      if (v === undefined) return this._randomizationFactor;
      this._randomizationFactor = v;
      (_a = this.backoff) === null || _a === void 0 ? void 0 : _a.setJitter(v);
      return this;
    }
  }, {
    key: "reconnectionDelayMax",
    value: function reconnectionDelayMax(v) {
      var _a;
      if (v === undefined) return this._reconnectionDelayMax;
      this._reconnectionDelayMax = v;
      (_a = this.backoff) === null || _a === void 0 ? void 0 : _a.setMax(v);
      return this;
    }
  }, {
    key: "timeout",
    value: function timeout(v) {
      if (!arguments.length) return this._timeout;
      this._timeout = v;
      return this;
    }
    /**
     * Starts trying to reconnect if reconnection is enabled and we have not
     * started reconnecting yet
     *
     * @private
     */
  }, {
    key: "maybeReconnectOnOpen",
    value: function maybeReconnectOnOpen() {
      // Only try to reconnect if it's the first time we're connecting
      if (!this._reconnecting && this._reconnection && this.backoff.attempts === 0) {
        // keeps reconnection from firing twice for the same reconnection loop
        this.reconnect();
      }
    }
    /**
     * Sets the current transport `socket`.
     *
     * @param {Function} fn - optional, callback
     * @return self
     * @public
     */
  }, {
    key: "open",
    value: function open(fn) {
      var _this2 = this;
      if (~this._readyState.indexOf("open")) return this;
      this.engine = new Socket$1(this.uri, this.opts);
      var socket = this.engine;
      var self = this;
      this._readyState = "opening";
      this.skipReconnect = false;
      // emit `open`
      var openSubDestroy = on(socket, "open", function () {
        self.onopen();
        fn && fn();
      });
      var onError = function onError(err) {
        _this2.cleanup();
        _this2._readyState = "closed";
        _this2.emitReserved("error", err);
        if (fn) {
          fn(err);
        } else {
          // Only do this if there is no fn to handle the error
          _this2.maybeReconnectOnOpen();
        }
      };
      // emit `error`
      var errorSub = on(socket, "error", onError);
      if (false !== this._timeout) {
        var timeout = this._timeout;
        // set timer
        var timer = this.setTimeoutFn(function () {
          openSubDestroy();
          onError(new Error("timeout"));
          socket.close();
        }, timeout);
        if (this.opts.autoUnref) {
          timer.unref();
        }
        this.subs.push(function () {
          _this2.clearTimeoutFn(timer);
        });
      }
      this.subs.push(openSubDestroy);
      this.subs.push(errorSub);
      return this;
    }
    /**
     * Alias for open()
     *
     * @return self
     * @public
     */
  }, {
    key: "connect",
    value: function connect(fn) {
      return this.open(fn);
    }
    /**
     * Called upon transport open.
     *
     * @private
     */
  }, {
    key: "onopen",
    value: function onopen() {
      // clear old subs
      this.cleanup();
      // mark as open
      this._readyState = "open";
      this.emitReserved("open");
      // add new subs
      var socket = this.engine;
      this.subs.push(on(socket, "ping", this.onping.bind(this)), on(socket, "data", this.ondata.bind(this)), on(socket, "error", this.onerror.bind(this)), on(socket, "close", this.onclose.bind(this)), on(this.decoder, "decoded", this.ondecoded.bind(this)));
    }
    /**
     * Called upon a ping.
     *
     * @private
     */
  }, {
    key: "onping",
    value: function onping() {
      this.emitReserved("ping");
    }
    /**
     * Called with data.
     *
     * @private
     */
  }, {
    key: "ondata",
    value: function ondata(data) {
      try {
        this.decoder.add(data);
      } catch (e) {
        this.onclose("parse error", e);
      }
    }
    /**
     * Called when parser fully decodes a packet.
     *
     * @private
     */
  }, {
    key: "ondecoded",
    value: function ondecoded(packet) {
      var _this3 = this;
      // the nextTick call prevents an exception in a user-provided event listener from triggering a disconnection due to a "parse error"
      nextTick(function () {
        _this3.emitReserved("packet", packet);
      }, this.setTimeoutFn);
    }
    /**
     * Called upon socket error.
     *
     * @private
     */
  }, {
    key: "onerror",
    value: function onerror(err) {
      this.emitReserved("error", err);
    }
    /**
     * Creates a new socket for the given `nsp`.
     *
     * @return {Socket}
     * @public
     */
  }, {
    key: "socket",
    value: function socket(nsp, opts) {
      var socket = this.nsps[nsp];
      if (!socket) {
        socket = new Socket(this, nsp, opts);
        this.nsps[nsp] = socket;
      } else if (this._autoConnect && !socket.active) {
        socket.connect();
      }
      return socket;
    }
    /**
     * Called upon a socket close.
     *
     * @param socket
     * @private
     */
  }, {
    key: "_destroy",
    value: function _destroy(socket) {
      var nsps = Object.keys(this.nsps);
      for (var _i = 0, _nsps = nsps; _i < _nsps.length; _i++) {
        var nsp = _nsps[_i];
        var _socket = this.nsps[nsp];
        if (_socket.active) {
          return;
        }
      }
      this._close();
    }
    /**
     * Writes a packet.
     *
     * @param packet
     * @private
     */
  }, {
    key: "_packet",
    value: function _packet(packet) {
      var encodedPackets = this.encoder.encode(packet);
      for (var i = 0; i < encodedPackets.length; i++) {
        this.engine.write(encodedPackets[i], packet.options);
      }
    }
    /**
     * Clean up transport subscriptions and packet buffer.
     *
     * @private
     */
  }, {
    key: "cleanup",
    value: function cleanup() {
      this.subs.forEach(function (subDestroy) {
        return subDestroy();
      });
      this.subs.length = 0;
      this.decoder.destroy();
    }
    /**
     * Close the current socket.
     *
     * @private
     */
  }, {
    key: "_close",
    value: function _close() {
      this.skipReconnect = true;
      this._reconnecting = false;
      this.onclose("forced close");
      if (this.engine) this.engine.close();
    }
    /**
     * Alias for close()
     *
     * @private
     */
  }, {
    key: "disconnect",
    value: function disconnect() {
      return this._close();
    }
    /**
     * Called upon engine close.
     *
     * @private
     */
  }, {
    key: "onclose",
    value: function onclose(reason, description) {
      this.cleanup();
      this.backoff.reset();
      this._readyState = "closed";
      this.emitReserved("close", reason, description);
      if (this._reconnection && !this.skipReconnect) {
        this.reconnect();
      }
    }
    /**
     * Attempt a reconnection.
     *
     * @private
     */
  }, {
    key: "reconnect",
    value: function reconnect() {
      var _this4 = this;
      if (this._reconnecting || this.skipReconnect) return this;
      var self = this;
      if (this.backoff.attempts >= this._reconnectionAttempts) {
        this.backoff.reset();
        this.emitReserved("reconnect_failed");
        this._reconnecting = false;
      } else {
        var delay = this.backoff.duration();
        this._reconnecting = true;
        var timer = this.setTimeoutFn(function () {
          if (self.skipReconnect) return;
          _this4.emitReserved("reconnect_attempt", self.backoff.attempts);
          // check again for the case socket closed in above events
          if (self.skipReconnect) return;
          self.open(function (err) {
            if (err) {
              self._reconnecting = false;
              self.reconnect();
              _this4.emitReserved("reconnect_error", err);
            } else {
              self.onreconnect();
            }
          });
        }, delay);
        if (this.opts.autoUnref) {
          timer.unref();
        }
        this.subs.push(function () {
          _this4.clearTimeoutFn(timer);
        });
      }
    }
    /**
     * Called upon successful reconnect.
     *
     * @private
     */
  }, {
    key: "onreconnect",
    value: function onreconnect() {
      var attempt = this.backoff.attempts;
      this._reconnecting = false;
      this.backoff.reset();
      this.emitReserved("reconnect", attempt);
    }
  }]);
}(Emitter);

/**
 * Managers cache.
 */
var cache = {};
function lookup(uri, opts) {
  if (_typeof(uri) === "object") {
    opts = uri;
    uri = undefined;
  }
  opts = opts || {};
  var parsed = url(uri, opts.path || "/socket.io");
  var source = parsed.source;
  var id = parsed.id;
  var path = parsed.path;
  var sameNamespace = cache[id] && path in cache[id]["nsps"];
  var newConnection = opts.forceNew || opts["force new connection"] || false === opts.multiplex || sameNamespace;
  var io;
  if (newConnection) {
    io = new Manager(source, opts);
  } else {
    if (!cache[id]) {
      cache[id] = new Manager(source, opts);
    }
    io = cache[id];
  }
  if (parsed.query && !opts.query) {
    opts.query = parsed.queryKey;
  }
  return io.socket(parsed.path, opts);
}
// so that "lookup" can be used both as a function (e.g. `io(...)`) and as a
// namespace (e.g. `io.connect(...)`), for backward compatibility
Object.assign(lookup, {
  Manager: Manager,
  Socket: Socket,
  io: lookup,
  connect: lookup
});

var InitModule = function InitModule(ctx, logger, nk, initializer) {
  initializer.registerRpc("typescript_healthcheck", rpcHealthcheck);
  var socket = lookup("ws://localhost:9999");
  socket.on("nft-transfer-observed", function (msg) {
    return logger.info(msg);
  });
  logger.info("Hello World!");
};
!InitModule && InitModule.bind(null);
