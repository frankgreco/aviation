// http://stepansuvorov.com/blog/2016/09/settimeout-es6-promise/
let timeout;

export function reset(ms, query){
    clearTimeout(timeout)

    let promiseCancel, promise = new Promise((resolve, reject) => {
        timeout = setTimeout(resolve(query), ms);
        promiseCancel = () => {
          clearTimeout(timeout); 
          reject(Error("Cancelled"));
        }
    });
    promise.cancel = () => { 
      promiseCancel();
    };
    return promise; 
}

export function allEnabled(searchFilters) {
  return Object.keys(searchFilters).filter(k => !searchFilters[k].enabled).length === 0
}

export function allDisabled(searchFilters) {
  return !allEnabled(searchFilters)
}
