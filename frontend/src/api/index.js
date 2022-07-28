import axios from 'axios'
// 引入进度条
import nprogress from "nprogress";
import "nprogress/nprogress.css";

const request = axios.create({
  baseURL: "http://127.0.0.1:4523/m1/1336335-0-default/api",
  timeout: 5000,
});

// 请求拦截器
request.interceptors.request.use(
  config => {
    nprogress.start()

    const token = localStorage.getItem("token")

    if (token && config.headers) config.headers.Authorization = token
    return config
  },
  err => Promise.reject(err)
)

// 响应拦截器
request.interceptors.response.use(
  rep => {
    nprogress.done()
    return rep.data
  },
  err => Promise.reject(err)
)

export default request