import React, { Fragment, useEffect, useState, useRef } from 'react'
// 引入 connect 用于连接 UI 组件和 redux
import { connect } from 'react-redux'
import { debounce } from '../../utils/throttle.js'
import './index.less'

function DirContent(props) {
  // console.log("DIRCONTENT GET: ", props);
  const { dirList, fileList } = props
  // console.log(dirList, fileList);

  const [crtClick, setCrtClick] = useState()
  const [crtFocus, setCrtFocus] = useState()
  const [menuBoxLeft, setMenuBoxLeft] = useState()
  const [menuBoxTop, setMenuBoxTop] = useState()
  const [showMenuBox, setShowMenuBox] = useState(false)

  const contextMenuRef = useRef()

  function getFileIcon(fileType = 'media') {
    switch (fileType) {
      case 'text':
        return (<svg t="1661408335321" className="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1687" width="16" height="16"><path d="M714 762.2h-98.2c-16.6 0-30 13.4-30 30s13.4 30 30 30H714c16.6 0 30-13.4 30-30s-13.4-30-30-30zM487.4 762.2H147.1c-16.6 0-30 13.4-30 30s13.4 30 30 30h340.3c16.6 0 30-13.4 30-30s-13.4-30-30-30z" fill="#33CC99" p-id="1688"></path><path d="M838.253 130.023l65.548 65.548-57.982 57.983-65.549-65.549z" fill="#FFB89A" p-id="1689"></path><path d="M743.7 955.9H195.8c-53.7 0-97.4-43.7-97.4-97.4V174.8c0-53.7 43.7-97.4 97.4-97.4H615c16.6 0 30 13.4 30 30s-13.4 30-30 30H195.8c-20.6 0-37.4 16.8-37.4 37.4v683.7c0 20.6 16.8 37.4 37.4 37.4h547.9c20.6 0 37.4-16.8 37.4-37.4v-395c0-16.6 13.4-30 30-30s30 13.4 30 30v395.1c0 53.6-43.7 97.3-97.4 97.3z" fill="#666" p-id="1690"></path><path d="M907.7 122.1l-39.2-39.2c-24-24-65.1-21.9-91.7 4.7L419.5 445 347 643.6l198.6-72.4L903 213.8c12.1-12.1 19.6-27.7 21.1-44 1.8-18.1-4.3-35.5-16.4-47.7zM512.6 519.3L447.5 543l23.7-65.1 264.7-264.7 40.9 41.7-264.2 264.4z m348-347.9l-41.3 41.3-40.9-41.7 40.9-40.9c3.1-3.1 6.2-3.9 7.6-3.9l37.6 37.6c-0.1 1.3-0.9 4.5-3.9 7.6z" fill="#ccc" p-id="1691"></path></svg>)
      case 'img':
        return (<svg t="1661408322144" className="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1529" width="16" height="16"><path d="M840.5 798.2L662.3 599.5l-151 173.7-173.7-173.7-167.7 201c-21 30.4 0.9 71.8 37.9 71.6l594.7-3.3c36.2-0.1 57.8-40.3 38-70.6z" fill="#FFB89A" p-id="1530"></path><path d="M741.6 647.3l-52.3-47.7c-12.2-11.2-31.2-10.3-42.4 1.9s-10.3 31.2 1.9 42.4l52.3 47.7c5.8 5.3 13 7.8 20.2 7.8 8.1 0 16.2-3.3 22.2-9.8 11.2-12.1 10.3-31.1-1.9-42.3zM631.2 546.5c-12.4-11-31.4-9.8-42.3 2.6l-98.8 111.7-171-165.7L87.9 724.7c-11.8 11.7-11.8 30.7-0.1 42.4 5.9 5.9 13.6 8.9 21.3 8.9 7.6 0 15.3-2.9 21.1-8.7l189.4-188.1 173.8 168.5L633.8 589c11-12.5 9.8-31.5-2.6-42.5z" fill="#33CC99" p-id="1531"></path><path d="M721.3 342.8m-35.1 0a35.1 35.1 0 1 0 70.2 0 35.1 35.1 0 1 0-70.2 0Z" fill="#33CC99" p-id="1532"></path><path d="M743.2 175.1H191.6c-70.6 0-128.3 57.7-128.3 128.3v499.2c0 70.6 57.7 128.3 128.3 128.3h551.5c70.6 0 128.3-57.7 128.3-128.3V303.5c0.1-70.6-57.7-128.4-128.2-128.4z m68.3 627.6c0 18.1-7.1 35.2-20.1 48.2-13 13-30.1 20.1-48.2 20.1H191.6c-18.1 0-35.2-7.1-48.2-20.1-13-13-20.1-30.1-20.1-48.2V303.5c0-18.1 7.1-35.2 20.1-48.2 13-13 30.1-20.1 48.2-20.1h551.5c18.1 0 35.2 7.1 48.2 20.1 13 13 20.1 30.1 20.1 48.2v499.2z" fill="#778" p-id="1533"></path><path d="M799.7 90.9H237.2c-16.6 0-30 13.4-30 30s13.4 30 30 30h562.4c26.1 0 50.8 10.3 69.4 28.9 18.6 18.6 28.9 43.3 28.9 69.4v482.4c0 16.6 13.4 30 30 30s30-13.4 30-30V249.2C958 161.9 887 90.9 799.7 90.9z" fill="skyblue" p-id="1534"></path></svg>)
      case 'media':
        return (<svg t="1661408294082" className="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1373" width="16" height="16"><path d="M861.9 383.8H218.1c-36.4 0-66.1-29.8-66.1-66.1V288c0-36.4 29.8-66.1 66.1-66.1h643.8c36.4 0 66.1 29.8 66.1 66.1v29.7c0 36.3-29.8 66.1-66.1 66.1z" fill="#FFB89A" p-id="1374"></path><path d="M822.9 129.2H199.8c-77.2 0-140.4 63.2-140.4 140.4v487.2c0 77.2 63.2 140.4 140.4 140.4h623.1c77.2 0 140.4-63.2 140.4-140.4V269.6c0-77.2-63.2-140.4-140.4-140.4z m80.4 177H760.4L864.6 201c5.4 3.3 10.4 7.3 15 11.8 15.3 15.3 23.7 35.4 23.7 56.8v36.6z m-673.3 0l104-117h61.3l-109.1 117H230z m247.4-117h169.2L532 306.2H368.3l109.1-117z m248.8 0h65.6L676 306.2h-60l112.5-114.8-2.3-2.2zM143 212.9c15.3-15.3 35.4-23.7 56.8-23.7h53.9l-104 117h-30.4v-36.5c0.1-21.4 8.5-41.5 23.7-56.8z m736.6 600.7c-15.3 15.3-35.4 23.7-56.8 23.7h-623c-21.3 0-41.5-8.4-56.8-23.7-15.3-15.3-23.7-35.4-23.7-56.8V366.2h783.9v390.6c0.1 21.3-8.3 41.5-23.6 56.8z" fill="#555" p-id="1375"></path><path d="M400.5 770.6V430.9L534.1 508c14.3 8.3 19.3 26.6 11 41-8.3 14.3-26.6 19.3-41 11l-43.6-25.2v131.8l114.1-65.9-7.5-4.3c-14.3-8.3-19.3-26.6-11-41 8.3-14.3 26.6-19.3 41-11l97.5 56.3-294.1 169.9z" fill="#33CC99" p-id="1376"></path></svg>)
      default:
        return <svg t="1661408432074" className="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="7386" width="16" height="16"><path d="M735.48 881.936H291.888c-54.448 0-98.584-44.144-98.584-98.576V241.168c0-54.448 44.136-98.584 98.584-98.584H608.88c8.536-1.608 17.6 0.4 24.192 7l192.264 192.288c4.28 4.264 6.768 9.608 7.576 15.16 0.752 2.32 1.112 4.808 1.112 7.368v418.96c0.024 54.424-44.072 98.552-98.496 98.576h-0.048z m-98.592-651.6v109.416h109.424L636.888 230.336zM784.784 389.04H612.264a24.664 24.664 0 0 1-24.656-24.648V191.872H291.888a49.288 49.288 0 0 0-49.288 49.288V783.36c0 27.2 22.056 49.264 49.272 49.264h443.584a49.296 49.296 0 0 0 49.304-49.264V389.04h0.024z" p-id="7387" fill="#ffffff"></path></svg>
    }

  }

  function toTag() {

  }


  useEffect(() => {
    // 去除选择文字
    document.addEventListener('selectstart', (e) => {
      e.preventDefault()
    })
    document.addEventListener('contextmenu', (e) => {
      e.preventDefault()
    })
    document.addEventListener('click', (e) => {
      setShowMenuBox(false)
      if (e.target.id === 'dC_main') setCrtClick('')
    })


    function menuBox(e) {
      setShowMenuBox(false)
      setTimeout(() => {
        let menuBox = contextMenuRef.current
        let parentLeft = menuBox.offsetParent.offsetLeft
        let parentTop = menuBox.offsetParent.offsetTop
        setMenuBoxLeft(e.x - parentLeft)
        setMenuBoxTop(e.y - parentTop)
        setTimeout(() => {
          setShowMenuBox(true)
        }, 140)
      }, 100)
    }

    // 修改右键模式
    document.addEventListener('contextmenu', debounce(menuBox, 500, false))
  }, [])

  function toggle_DC_click(id) {
    setCrtClick(id)
  }

  function toggle_DC_focus(id) {
    setCrtFocus(id)
  }

  const menuList = ['重命名', '打开', '删除', '复制', '粘贴', '属性']


  const list = ['11', '23', '234']

  return (
    <div id="dirContent">
      <div id="dC_nav">{
        list.map((i, index) => {
          return (
            <Fragment>
              <span className='dc_nav_tap' onClick={toTag} title={'backTo: ' + i}>{i}</span>
              {index === list.length - 1 ? <span></span> : <span>&gt;</span>}
            </Fragment>
          )
        })
      }</div>
      <div id="dC_main">
        {
          dirList?.map((i, index) => {
            return (<div className={`dC_dir ${crtFocus === i._id && 'dC_focus'} ${crtClick === i._id && 'dC_click'} `}
              key={i._id + index}
              onClick={() => toggle_DC_click(i._id)}
              onContextMenu={() => toggle_DC_click(i._id)}
              onMouseEnter={() => toggle_DC_focus(i._id)}
              onMouseLeave={() => toggle_DC_focus('')}
            >
              <div className="dirIcon">
                <svg t="1661406678205" className="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1737" width="16" height="16"><path d="M439.9 359.2l-99.7-156.3h-81.4c-66.5 0-120.5 53.9-120.5 120.5v294.9h751.4V359.2H439.9z" fill="#2867CE" p-id="1738"></path><path d="M769.3 871.6H258.7c-92.9 0-168.5-75.6-168.5-168.5V323.3c0-92.9 75.6-168.5 168.5-168.5h81.4c16.4 0 31.7 8.4 40.5 22.2l85.5 134.1h423.5c26.6 0 48.1 21.5 48.1 48.1v343.9c0.1 92.9-75.5 168.5-168.4 168.5zM258.7 250.9c-39.9 0-72.4 32.5-72.4 72.4v379.8c0 39.9 32.5 72.4 72.4 72.4h510.5c39.9 0 72.4-32.5 72.4-72.4V407.3H439.9c-16.4 0-31.7-8.4-40.5-22.2L313.9 251h-55.2z" fill="#BDD2EF" p-id="1739"></path><path d="M840.2 272.7h-314l-69.6-109.2h383.6z" fill="#2867CE" p-id="1740"></path></svg>
              </div>
              <span className='dirName'>{i.name}</span>
            </div>)
          })
        }
        {
          fileList?.map((i, index) => {
            return (<div className={`dC_file ${crtClick === /*i._id*/index && 'dC_click'} ${crtFocus === index && 'dC_focus'}`}
              key={index}
              onClick={() => toggle_DC_click(/*i._id*/index)}
              onContextMenu={() => toggle_DC_click(index)}
              onMouseEnter={() => toggle_DC_focus(index)}
              onMouseLeave={() => toggle_DC_focus('')}>
              <div className="fileIcon">
                {getFileIcon()}
              </div>
              <span className='fileName'>{i}</span>
            </div>)
          })
        }
      </div>
      <div id="contextMenu" ref={contextMenuRef}
        style={{
          transform: `translate(${menuBoxLeft}px,${menuBoxTop}px)`,
          opacity: showMenuBox ? 1 : 0
        }}>
        {menuList.map((i, index) => {
          return (
            <span className='menuItem'
              key={index}
            >{i}</span>
          )
        })}
      </div>
    </div>
  )
}

// 暴露容器
export default connect(
  // 映射状态
  state => ({
    dirList: state.dirContent.dirList,
    fileList: state.dirContent.fileList
  }), {}
)(DirContent)
