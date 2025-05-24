<script setup>
import { reactive } from 'vue'
import { Gettestjson, Flashtime, Greet } from '../../wailsjs/go/service/App'
import { ElMessage } from 'element-plus'
import { EventsOn } from '../../wailsjs/runtime'

const data = reactive({
  name: '',
  resultText: '数据显示区域',
  systime: 'Loading...',
})

Flashtime()

EventsOn('time', (time) => {
  data.systime = time
})

function greet() {
  Greet(data.name).then((result) => {
    data.resultText = result
  })
}

function toast() {
  Greet(data.name)
    .then((result) => {
      ElMessage.success(result)
    })
    .catch((err) => {
      ElMessage.error(err)
    })
}

function convertBytes(byteSize) {
  byteSize = Number(byteSize)
  if (byteSize < 0) {
    throw new Error('Byte size cannot be negative')
  }

  const units = ['B', 'KB', 'MB', 'GB', 'TB']
  let index = 0

  while (byteSize >= 1024 && index < units.length - 1) {
    byteSize /= 1024
    index++
  }

  return `${byteSize.toFixed(2)} ${units[index]}`
}

function openDir() {
  ElMessage.success('获取文件中')
  Gettestjson()
    .then((result) => {
      const fileListDiv = document.getElementById('filelist')
      fileListDiv.innerHTML = ''
      const fileListinfo = document.getElementById('fileinfo')
      fileListinfo.innerHTML = ''

      const jsondata = JSON.parse(result)
      jsondata.data.list.forEach((item, index) => {
        const folderName = item.server_filename
        const folderPath = item.path
        const creationTime = new Date(item.local_ctime * 1000).toLocaleString()
        let fileSize = item.size
        let fileSizeFomat = convertBytes(fileSize)
        const fileCate = item.category
        const fsid = item.fs_id
        let fileCateText = '其他'

        switch (fileCate) {
          case 1:
            fileCateText = '视频'
            break
          case 2:
            fileCateText = '音乐'
            break
          case 3:
            fileCateText = '图片'
            break
          case 4:
            fileCateText = '文档'
            break
          case 5:
            fileCateText = '应用'
            break
          case 6:
            fileCateText = '其他'
            break
          case 7:
            fileCateText = '种子'
            break
          default:
            fileCateText = '未知类型'
            break
        }

        let clickCommend = `javascript:openMeue('${index + 1}');`
        const isDir = item.isdir
        let MenuText = '菜单'

        if (isDir === 1) {
          fileSizeFomat = '---'
          fileCateText = '文件夹'
          clickCommend = `javascript:openDir('${encodeURIComponent(folderPath)}');`
          MenuText = '文件夹'
        }

        const html = `
          <tr class="el-table__row">
            <td colspan="1" rowspan="1" class="el-table__cell">
              <div class="cell">${folderName}</div>
            </td>
            <td colspan="1" rowspan="1" class="el-table__cell">
              <div class="cell">${fileCateText}</div>
            </td>
            <td colspan="1" rowspan="1" class="el-table__cell">
              <div class="cell">${fileSizeFomat}</div>
            </td>
            <td colspan="1" rowspan="1" class="el-table__cell">
              <div class="cell">${creationTime}</div>
            </td>
            <td colspan="1" rowspan="1" class="el-table-fixed-column--right is-first-column el-table__cell">
              <div class="cell">
                <a aria-labelledby="js_p1m1_bd" href="${clickCommend}" class="el-button el-button--primary el-button--small is-link">打开${MenuText}</a>
              </div>
            </td>
          </tr>
        `

        const info = `
          <div id="Info_${index + 1}" hidden>
            <div id="File_Type_${index + 1}" value="${fileCate}"></div>
            <div id="File_Path_${index + 1}" value="${folderPath}"></div>
            <div id="File_Fsid_${index + 1}" value="${fsid}"></div>
            <div id="File_Size_${index + 1}" value="${fileSize}"></div>
            <div id="File_Filename_${index + 1}" value="${folderName}"></div>
          </div>
        `

        fileListDiv.innerHTML += html
        fileListinfo.innerHTML += info
      })
    })
    .catch((err) => {
      ElMessage.error(err)
    })
}
</script>

<template>
  <div class="inner" style="overflow:auto; height: 100vh;">
    <h1>前后端通信演示</h1>
    <div id="result" class="result">{{ data.resultText }}</div>

    <!-- 输入框和按钮 -->
    <div style="display: flex; align-items: center; gap: 10px; margin-bottom: 20px;">
      <el-input v-model="data.name" placeholder="输入内容" style="flex: 1;" />
      <el-button @click="greet" type="primary" style="flex: 0 0 auto;">Greet</el-button>
      <el-button @click="toast" type="primary" style="flex: 0 0 auto;">Toast</el-button>
      <el-button @click="openDir" type="primary" style="flex: 0 0 auto;">forEach</el-button>
    </div>

    <!-- 时间 -->
    <div class="welcome">
      <p>当前时间：<span class="datetime">{{ data.systime }}</span></p>
    </div>
    <!-- 表格 -->
    <div id="all">
      <div class="el-pagination is-background el-header">
        <ul class="el-pager" id="tablab">
          <li class="number" aria-current="true" aria-label="page 1" tabindex="0">
            <a
              href="javascript:openDir('%2F');"
              role="button"
              class="el-tabs__item is-top is-active"
              title="全部文件"
              >全部文件</a
            >
          </li>
        </ul>
      </div>

      <div
        class="el-main el-table--fit el-table--scrollable-x el-table--enable-row-transition el-table el-table--layout-fixed is-scrolling-left"
        data-prefix="el"
        style="width: 100%; overflow-x: auto;"
      >
        <div class="el-table__header-wrapper">
          <table class="el-table__header" border="0" cellpadding="0" cellspacing="0" style="width: 100%;">
            <thead>
              <tr>
                <th class="is-leaf el-table__cell" colspan="1" rowspan="1" style="min-width: 150px;">
                  <div class="cell">文件名</div>
                </th>
                <th class="is-leaf el-table__cell" colspan="1" rowspan="1" style="min-width: 100px;">
                  <div class="cell">文件类型</div>
                </th>
                <th class="is-leaf el-table__cell" colspan="1" rowspan="1" style="min-width: 100px;">
                  <div class="cell">文件大小</div>
                </th>
                <th class="is-leaf el-table__cell" colspan="1" rowspan="1" style="min-width: 150px;">
                  <div class="cell">创建时间</div>
                </th>
                <th class="el-table_13_column_42 el-table-fixed-column--right is-first-column is-leaf el-table__cell" colspan="1" rowspan="1" style="min-width: 100px;">
                  <div class="cell">操作</div>
                </th>
              </tr>
            </thead>
          </table>
        </div>
        <div class="el-table__body-wrapper" style="overflow-y: auto; max-height: 400px;">
          <div class="el-scrollbar">
            <div class="el-scrollbar__wrap el-scrollbar__wrap--hidden-default">
              <div class="el-scrollbar__view" style="display: inline-block; vertical-align: middle;">
                <table class="el-table__body" cellspacing="0" cellpadding="0" border="0" style="width: 100%;">
                  <tbody tabindex="-1" id="filelist"></tbody>
                </table>
                <div id="fileinfo"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* 基础样式 */
.result {
  margin-bottom: 20px;
}

.welcome {
  text-align: center;
  margin: 20px 0;
}

#all {
  width: 100%;
  max-width: 1200px; /* 限制最大宽度，避免内容过宽 */
  margin: 0 auto; /* 居中对齐 */
}

.el-table__header th,
.el-table__body td {
  padding: 8px;
}

/* 确保表格容器支持滚动 */
.inner {
  overflow: auto;
  height: 100vh;
}

.el-table__body-wrapper {
  overflow-y: auto;
  max-height: 400px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .el-table__header th,
  .el-table__body td {
    font-size: 12px;
    padding: 5px;
  }

  .el-button {
    font-size: 12px;
    padding: 5px 10px;
  }
}

@media (max-width: 480px) {
  .el-table__header th,
  .el-table__body td {
    font-size: 10px;
    padding: 3px;
  }

  .el-button {
    font-size: 10px;
    padding: 3px 8px;
  }
}
</style>