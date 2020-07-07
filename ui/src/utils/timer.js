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