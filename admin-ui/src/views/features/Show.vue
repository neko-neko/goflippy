<template>
  <div>
    <h1>{{feature.name}}</h1>
    <p>last_activated_at: {{feature.last_activated_at}}</p>

    <h2>ストラテジ</h2>
    <div>
      <div>
        <span>simple</span>

        <label>On</label>
        <input type="radio">

        <label>Off</label>
        <input type="radio">
      </div>

      <div>
        <label>group</label>
        <input type="text">
      </div>

      <div>
        <label>attribute</label>
        <input type="text">
      </div>

      <div>
        <label>release_date_time</label>
        <input type="date">
      </div>

      <div>
        <label>percentage</label>
        <input type="number">
      </div>
    </div>

    <!-- <ul v-if="feature.filters.length > 0">
      <li></li>
    </ul>
    <p v-else>フィルタは設定されていません</p> -->
    <v-btn color="primary">Save</v-btn>
  </div>
</template>

<script>
import { apiUrl, fetchData, formatDate } from "../../utils"

export default {
  data: () => ({
    feature: {
      filters: {
        group: {
          enable: false,
          value: ''
        },
        attribute: {
          enable: false,
          value: ''
        },
        uuid: {
          enable: false,
          value: ''
        },
        release_date_time: {
          enable: false,
          value: ''
        },
        percentage: {
          enable: false,
          value: ''
        }
      }
    }
  }),
  created() {
    const self = this
    const projectId = self.$route.params.id
    const apiKey = self.$route.params.key
    const url = apiUrl.getFeatureUrl(projectId,apiKey)

    fetchData(url).then(data => {
      const feature = data.feature

      if (feature) {
        feature.last_activated_at = formatDate(feature.last_activated_at)
        self.feature = feature
      }
    })
  }
}
</script>
