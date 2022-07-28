import routes from './router'
import './App.less'
import { useRoutes } from 'react-router-dom'
import 'animate.css'

export default function App() {
  const element = useRoutes(routes)

  return (
    <div id="app">
      {element}
    </div>
  )
}

