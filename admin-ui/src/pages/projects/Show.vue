<template>
  <div>
    <h1>{{project.name}}</h1>
    <p>{{project.description}}</p>
    <h2>{{ $t('features.features') }}</h2>
    <div>
      <div>
        <v-text-field :label="$t('features.name')" v-model="featureForm.name"></v-text-field>
      </div>
      <div>
        <label>{{ $t('features.key') }}</label>
        <input v-model="featureForm.key" :placeholder="$t('features.key')">
      </div>
      <v-btn @click="postFeature">{{ $t('create') }}</v-btn>
    </div>
    <ul v-if="features.length > 0">
      <li v-for="feature in features" :key="feature.key">
        <router-link :to="`/projects/${project._id}/features/${feature.key}`">{{feature.name}}</router-link>
      </li>
    </ul>
    <p v-else>{{ $t('projects.error.no_features') }}</p>
    <h2>{{ $t('users.users') }}</h2>
    <ul v-if="users.length > 0">
      <li v-for="user in users" :key="user.uuid">
        <p>
          {{user.first_name}}&nbsp;{{user.last_name}}
          <router-link :to="`/projects/${project._id}/users/${user.uuid}`">{{ $t('detail') }}</router-link>
        </p>
      </li>
    </ul>
    <p v-else>{{ $t('projects.error.no_users') }}</p>
  </div>
</template>

<script>
import { apiUrl, fetchData, postData } from "../../utils";

export default {
  data: () => ({
    project: {
      id: "",
      name: "",
      description: ""
    },
    features: [],
    users: [],
    featureForm: {
      name: "",
      key: ""
    }
  }),
  created() {
    const self = this;
    const projectId = self.$route.params.id;
    const getProjectUrl = apiUrl.getProjectUrl(projectId);
    const getFeaturesUrl = apiUrl.getFeaturesUrl(projectId);
    const getUsersUrl = apiUrl.getUsersUrl(projectId);

    fetchData(getProjectUrl).then(data => {
      const project = data.project;

      if (project) self.project = project;
    });

    fetchData(getFeaturesUrl).then(data => {
      const features = data.features;

      if (features) self.features = features;
    });

    fetchData(getUsersUrl).then(data => {
      const users = data.users;

      if (users) self.users = users;
    });
  },
  methods: {
    postFeature() {
      const self = this;
      const projectId = self.$route.params.id;
      const url = apiUrl.postFeatureUrl(projectId);

      const params = {
        name: self.featureForm.name,
        key: self.featureForm.key
      };

      postData(url, params).then(data => {
        const feature = data.feature;

        if (feature) self.features.push(feature);
      });
    }
  }
};
</script>
