// 引入 type 类型常量
import { CHANGE } from '../constant'

// 定义一个初始值
const initState = { dirList: [], fileList: [] }

// 暴露一个为 Count 组件服务的 reducer
// reducer 是一个纯函数
export default function countReducer(preState = initState, action) {
  // 从 action 对象中获取 type data
  const { type, data } = action
  // 根据 type 决定如何加工数据
  if (type === CHANGE) {
    return data
  } else {
    return preState
  }
}