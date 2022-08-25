// 专门为 Count 组件生成 action 对象
import { CHANGE } from '../constant'

// 返回体为对象的 同步action
export const change = data => ({ type: CHANGE, data })