export const throttle = function (fn) {
  let ff = fn;
  return function () {
    if (fn) {
      fn.apply(this, arguments);
      fn = null;
      setTimeout(() => {
        fn = ff;
      }, 400);
    }
  };
};

export const debounce = (fn, delay = 500, immediate = true) => {
  let timer
  // 保留初始值, tmp 代表该函数是否需要立刻执行
  let tmp = immediate

  const debounced = (...args) => {

    if (immediate) {
      fn.apply(this, args);
      immediate = false
    } else {
      clearTimeout(timer);
    }

    timer = window.setTimeout(() => {
      // 恢复
      immediate = tmp
      // 返回可能存在的返回值
      return tmp ? null : fn.apply(this, args);
    }, delay)
  }

  return debounced
}