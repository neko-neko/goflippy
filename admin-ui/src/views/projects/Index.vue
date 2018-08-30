<template>
  <v-flex xs10 offset-xs1>
    <h1 class="headline">プロジェクト一覧</h1>
    <v-dialog v-model="dialog" max-width="500px">
      <v-card>
        <v-card-title>
          <span class="headline">プロジェクト追加</span>
        </v-card-title>
        <v-card-text>
          <v-text-field box label="名前"
            v-model="projectForm.name"
          />

          <v-textarea box label="説明"
            v-model="projectForm.description"
          />
        </v-card-text>
        <v-card-actions>
          <v-btn @click="postProject" color="primary">作成</v-btn>
          <v-btn color="blue darken-1" outline @click.native="dialog = false">閉じる</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-layout justify-end row>
      <v-btn @click="dialog = true" color="primary">プロジェクト追加</v-btn>
    </v-layout>

    <v-card>
      <v-data-table
        id="projects"
        hide-actions
        :headers="headers"
        :items="projects"
      >
        <template slot="items" slot-scope="props">
          <tr>
            <td>{{props.item.name}}</td>
            <td>{{props.item.description}}</td>
            <td><router-link :to="`/projects/${props.item._id}`">詳細</router-link></td>
          </tr>
        </template>
      </v-data-table>
    </v-card>
  </v-flex>
</template>

<script>
import { apiUrl, fetchData, postData } from "../../utils"

export default {
  data: () => ({
    projectForm: {
      name: '',
      description: '',
    },
    dialog: false,
    headers: [
      { text: 'プロジェクト名', value: 'name' },
      { text: '説明', value: 'description' },
      { text: '', value: '' },
    ],
    projects: [],
  }),
  created() {
    const self = this
    const url = apiUrl.getProjectsUrl()

    fetchData(url).then(data => {
      const projects = data.projects

      if (projects) self.projects = projects
    })
  },
  methods: {
    postProject() {
      const self = this
      const url = apiUrl.postProjectUrl()

      const params = {
        name: self.projectForm.name,
        description: self.projectForm.description,
      }

      postData(url, params).then(data => {
        const project = data.project

        if (project) self.projects.push(project)
        self.clearForm()
        self.dialog = false
      })
    },
    clearForm() {
      this.projectForm.name = ''
      this.projectForm.description = ''
    }
  }
}
</script>
