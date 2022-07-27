import { useRoutes } from 'react-router-dom'
import routes from './routes'
import './App.less'

function App() {

  const element = useRoutes(routes)

  return (
    <div id="app">
      {/* <Home/> */}
    </div>
  )
}

export default App
