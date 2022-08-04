import React, {useState, useRef, useEffect} from 'react'
import { useNavigate } from 'react-router-dom'

import Footer from '../../container/Footer'
import request from '../../api'
import './index.less'

// 读取配置
import settings from '../../assets/settings.json'

export default function Login() {
  const [logging, setLogging] = useState(false)

  const Password = useRef()
  const Image = useRef()
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
        setTimeout(() => navigate('/home', {replace:true}),500)
      }else{
        Password.current.value = ''
        alert("no no no")
      }
      setTimeout(() => setLogging(false),500)
    })
  }

  // 遗忘密码
  const forget = () => {alert("我傻了, 你密码我怎么知道")}

  return (
  <div id="login">
    <img className="bgImg" ref={Image} src={settings.Login.backgroundImagePath} alt="出错了呢" />
    <div className="loginBox animate__animated animate__fadeIn">
      <div className='password'>
        <input type="password" ref={Password} placeholder="password"/>
      </div>
      <button onClick={login}>
        {logging ? <span className="iconfont icon-jiazai_shuang loging"></span> : <span className="sign_in">Sign In</span>}
      </button>
        <span className="forget" onClick={forget}>Forget password ?</span>
    </div>
    <Footer/>
  </div>
  )
}
