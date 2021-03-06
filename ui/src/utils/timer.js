// http://stepansuvorov.com/blog/2016/09/settimeout-es6-promise/
let timeout;

export function reset(ms, query) {
  clearTimeout(timeout);

  let promiseCancel;

  const promise = new Promise((resolve, reject) => {
    timeout = setTimeout(resolve(query), ms);
    promiseCancel = () => {
      clearTimeout(timeout);
      reject(Error('Cancelled'));
    };
  });
  promise.cancel = () => {
    promiseCancel();
  };
  return promise;
}

export function buildQuery(searchFilters) {
  const filters = [];

  Object.keys(searchFilters).forEach((k) => {
    if (searchFilters[k] !== undefined && searchFilters[k].value !== undefined && searchFilters[k].value !== '') {
      filters.push(`${searchFilters[k].key}="${searchFilters[k].value}"`);
    }
  });

  return filters.join(' AND ');
}
