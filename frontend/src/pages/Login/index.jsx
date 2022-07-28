import React, {useState, useRef} from 'react'
import { useNavigate } from 'react-router-dom'
import request from '../../api'
import './index.less'

export default function Login() {
  const [logging, setLogging] = useState(false)

  const Password = useRef()
  const navigate = useNavigate()

  const login = () => {
    setLogging(true)
    let password = Password.current.value

    request({
      url:'/account',
      method: 'post',
      data: {
        password
      }
    }).then(response => {
      const {status, data:{token}} = response
      if(response.status === 200 && token){
        localStorage.setItem("token", "Bearer " + token)
        Password.current.value = ''
        // setTimeout(() => navigate('/home', {replace:true}),500)
      }else{
        alert("no no no")
      }
      setTimeout(() => setLogging(false),500)
      
    })

  }

  const forget = () => {}

  return (
    <div id="login">
      <div className="loginBox animate__animated animate__fadeIn">
        <div className="title">Anime Cat</div>
        <div className='password'>
          <input type="password" ref={Password} placeholder="password"/>
        </div>
        <button onClick={login}>
          {logging ? <span className="iconfont icon-jiazai_shuang loging"></span> : <span className="sign_in">Sign In</span>}
        </button>
          <span className="forget" onClick={forget}>Forget password ?</span>
      </div>
    </div>
  )
}
