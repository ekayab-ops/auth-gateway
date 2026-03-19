const { URL } = require('url');
const { parse } = require('querystring');

const parseUrl = (url) => {
  const parsedUrl = new URL(url);
  const queryParams = parse(parsedUrl.search);

  return {
    protocol: parsedUrl.protocol.replace(':', ''),
    hostname: parsedUrl.hostname,
    port: parsedUrl.port,
    path: parsedUrl.pathname,
    query: queryParams,
    fragment: parsedUrl.hash.replace('#', '')
  };
};

module.exports = {
  parseUrl
};