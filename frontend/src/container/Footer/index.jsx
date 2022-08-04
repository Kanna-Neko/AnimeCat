import React from 'react'
import './index.less'

export default function Footer() {
  const record = '浙ICP备xxxxx号'
  const copyRight = 'Copyright © 2022'

  return (
    <div id="footer">
      <p>{record} {copyRight}</p>
    </div>
  )
}
