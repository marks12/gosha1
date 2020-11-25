<template>
  <WorkSpace width="fit" height="fit" :class="{'bg-filter' : entityItem.IsFilter}">
    <VSet divider vertical height="fit">
      <VSet v-if="!entityItem.IsFilter" class="head-row">
        <VSpoiler
            :opened="spoilerStateOpened"
            @change-spoiler-state="spoilerStateOpened = !spoilerStateOpened"
        >
          <template #head>
            <VHead>{{ entityItem.Name }}</VHead>
          </template>
          <template #preview>
              <VSet vertical has-no-indent style="max-width: 220px;">
                <VText color="weak" v-if="entityItem.CommentType">
                  Types: {{ entityItem.CommentType | substring(40) }}
                </VText>
                <VText color="weak" v-if="entityItem.CommentDb">
                  Db: {{ entityItem.CommentDb | substring(40) }}
                </VText>
              </VSet>
          </template>
          <template #content>
            <VText style="max-width: 240px;">
              <VSet vertical  has-no-indent>
                <VText color="weak" v-if="entityItem.CommentType">
                  Types: {{ entityItem.CommentType }}
                </VText>
                <VText color="weak" v-if="entityItem.CommentDb">
                  Db: {{ entityItem.CommentDb }}
                </VText>
              </VSet>
            </VText>
          </template>
        </VSpoiler>
        <VSign width="M">Db</VSign>
        <VSign width="M">Types</VSign>
      </VSet>
      <VSet v-else vertical>
        <VSpoiler
            v-if="entityItem.CommentType"
            :opened="spoilerStateOpened"
            @change-spoiler-state="spoilerStateOpened = !spoilerStateOpened"
        >
          <template #head>
            <VHead>{{ entityItem.Name }}</VHead>
          </template>
          <template #preview>
            <VSet vertical has-no-indent>
              <VText color="weak" v-if="entityItem.CommentType">
                {{ entityItem.CommentType | substring(40) }}
              </VText>
            </VSet>
          </template>
          <template #content>
            <VText style="max-width: 240px;">
              <VSet vertical  has-no-indent>
                <VText color="weak" v-if="entityItem.CommentType">
                  {{ entityItem.CommentType }}
                </VText>
              </VSet>
            </VText>
          </template>
        </VSpoiler>
        <VSet v-else class="head-row">
          <VText :title="entityItem.Name">{{ entityItem.Name }}</VText>
          <VSign width="M">Type</VSign>
        </VSet>

      </VSet>

      <VSet v-if="!entityItem.IsFilter">
        <VButton @click="reqFind"
                 small :title="'Send Find Request'"
                 :class="entityItem.HttpMethods.IsFind ? 'small-button find' : 'small-button weak'" text="F"></VButton>

        <VButton @click="reqCreate"
                 small :title="'Send Create Request'"
                 :class="entityItem.HttpMethods.IsCreate ? 'small-button create' : 'small-button weak'" text="C">C
        </VButton>

        <VButton @click="reqRead"
                 small :title="'Send Read Request'"
                 :class="entityItem.HttpMethods.IsRead ? 'small-button read' : 'small-button weak'" text="R">R
        </VButton>

        <VButton @click="reqUpdate"
                 small :title="'Send Update Request'"
                 :class="entityItem.HttpMethods.IsUpdate ? 'small-button update' : 'small-button weak'" text="U">U
        </VButton>

        <VButton @click="reqDelete"
                 small :title="'Send Delete Request'"
                 :class="entityItem.HttpMethods.IsDelete ? 'small-button delete' : 'small-button weak'" text="D">D
        </VButton>

        <VButton @click="reqFoC"
                 small :title="'Send Find or create Request'"
                 :class="entityItem.HttpMethods.IsFindOrCreate ? 'small-button foc' : 'small-button weak'" text="FoC"
                 title="Find or create">FindOrCreate
        </VButton>

        <VButton @click="reqUoC"
                 small :title="'Send Update or create Request'"
                 :class="entityItem.HttpMethods.IsUpdateOrCreate ? 'small-button uoc' : 'small-button weak'" text="UoC"
                 title="Update or create">UoC
        </VButton>

      </VSet>


      <VSet vertical v-if="!entityItem.IsFilter">
        <template v-for="(field, i) in entityItem.Fields">
          <VSpoiler v-if="field.CommentType || field.CommentDb" style="margin-top: 0; margin-bottom: 3px;" content-indent="XS">
            <template #head>
              <VSet class="field-row">
                <VText>
                  {{ field.Name }}
                </VText>
                <VText width="M">
                  <VIcon name="check" v-if="field.IsDb"></VIcon>
                  <VIcon name="minus" v-else></VIcon>
                </VText>
                <VText width="M">
                  <VIcon name="check" v-if="field.IsType"></VIcon>
                  <VIcon name="minus" v-else></VIcon>
                </VText>
              </VSet>
            </template>
            <template #content>
              <VSet vertical indent-size="XS" style="width: 320px; ">
                <VText color="weak">
                  <template v-if="field.CommentType && field.CommentDb">
                    Type:
                  </template>
                  {{ field.CommentType }}
                </VText>
                <VText color="weak">
                  <template v-if="field.CommentDb">
                    Db:
                  </template>
                  {{ field.CommentDb }}
                </VText>
              </VSet>
            </template>
          </VSpoiler>
          <VSet v-else class="field-row" style="padding-left: 24px;">
            <VText>
              {{ field.Name }}
            </VText>
            <VText width="M">
              <VIcon name="check" v-if="field.IsDb"></VIcon>
              <VIcon name="minus" v-else></VIcon>
            </VText>
            <VText width="M">
              <VIcon name="check" v-if="field.IsType"></VIcon>
              <VIcon name="minus" v-else></VIcon>
            </VText>
          </VSet>

        </template>
      </VSet>
      <VSet vertical v-else>
        <template v-for="(field, i) in entityItem.Fields">

          <VSpoiler v-if="field.CommentType || field.CommentDb" style="margin-top: 0; margin-bottom: 3px;" content-indent="XS">
            <template #head>
              <VSet class="field-row">
                <VText>
                  {{ field.Name }}
                </VText>
                <VSign width="M">
                  {{ field.Type }}
                </VSign>
              </VSet>
            </template>
            <template #content>
              <VSet vertical indent-size="XS" style="width: 250px; ">
                <VText color="weak">
                  {{ field.CommentType }}
                </VText>
              </VSet>
            </template>
          </VSpoiler>

          <VSet v-else class="field-row" style="padding-left: 24px;">
            <VText>
              {{ field.Name }}
            </VText>
            <VSign width="M">
              {{ field.Type }}
            </VSign>
          </VSet>

        </template>
      </VSet>
    </VSet>
    <template slot="footer">
      <VSet>
        <VButton text="Edit" small @click="onEdit(entityItem)"></VButton>
      </VSet>
    </template>
  </WorkSpace>
</template>

<script>
import VBadge from "swtui/src/components/VBadge";
import VButton from "swtui/src/components/VButton";
import VSpoiler from "swtui/src/components/VSpoiler";
import VGroup from "swtui/src/components/VGroup";
import VSet from "swtui/src/components/VSet";
import VText from "swtui/src/components/VText";
import VHead from "swtui/src/components/VHead";
import VIcon from "swtui/src/components/VIcon";
import VSign from "swtui/src/components/VSign";
import WorkSpace from "swtui/src/components/WorkSpace";

export default {
  name: "EntityItem",
  components: {WorkSpace, VSpoiler, VBadge, VGroup, VSet, VText, VIcon, VSign, VButton, VHead},
  props: {
    entityItem: Object,
    onEdit: Function,
  },
  data() {
    return {
      spoilerStateOpened: false,
    };
  },
  methods: {
    onRequest(item) {
      this.$emit("onRequest", item);
    },
    reqFind() {
      this.entityItem.Action = 'find';
      this.onRequest(this.entityItem);
    },
    reqCreate() {
      this.entityItem.Action = 'create';
      this.onRequest(this.entityItem);
    },
    reqRead() {
      this.entityItem.Action = 'read';
      this.onRequest(this.entityItem);
    },
    reqUpdate() {
      this.entityItem.Action = 'update';
      this.onRequest(this.entityItem);
    },
    reqDelete() {
      this.entityItem.Action = 'delete';
      this.onRequest(this.entityItem);
    },
    reqFoC() {
      this.entityItem.Action = 'findorcreate';
      this.onRequest(this.entityItem);
    },
    reqUoC() {
      this.entityItem.Action = 'updateorcreate';
      this.onRequest(this.entityItem);
    },
  },
}
</script>

<style scoped lang="scss">

@import "~swtui/src/scss/variables";

.bg-filter {
  background: #daffc9;
}

.sw-icon {
  background: transparent;
}

.sw-button_small.find {
  background: $color-action;
  color: $color-bg-light;
}

.sw-button_small.create {
  background: $color-attention-secondary;
  color: $color-bg-light;
}

.sw-button_small.read {
  background: $color-action;
  color: $color-bg-light;
}

.sw-button_small.update {
  background: $color-selection-light;
  color: $color-bg-light;
}

.sw-button_small.delete {
  background: $color-attention;
  color: $color-bg-light;
}

.small-button {
  padding: 0 12px;
  margin: 0 5px;
}

.field-row {
  margin-bottom: 0 !important;
  margin-top: 0 !important;
  padding-right: 30px;
}

.head-row {
  padding-right: 30px;
}

</style>