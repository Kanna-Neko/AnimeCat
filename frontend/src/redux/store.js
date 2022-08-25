// applyMiddleware 应用中间件
import { createStore } from 'redux'

// 引入汇总后的 reducer
import reducer from './reducer'

// 暴露 store 对象
export default createStore(reducer)