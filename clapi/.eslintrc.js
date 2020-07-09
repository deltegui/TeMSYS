module.exports = {
    "env": {
        "es6": true,
        "node": true
    },
    "extends": "airbnb",
    "globals": {
        "Atomics": "readonly",
        "SharedArrayBuffer": "readonly"
    },
    "parserOptions": {
        "ecmaVersion": 2018,
        "sourceType": "module"
    },
    "rules": {
        'arrow-parens': 0,
        'object-curly-newline': 0,
        'no-confusing-arrow': 0,
        'keyword-spacing': 0,
        'quotes': 0,
        'no-plusplus': 0,
        'class-methods-use-this': 0,
        'linebreak-style': 0,
    }
};