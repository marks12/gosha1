package cmd

const usualEntityVueComponentData = `
	import {{Entity}, {Entity}Filter} from '../apiModel';

let {entity}Data = function () {
    return {
        {entity}Filter: new {Entity}Filter(),
        panel: {
            type: '',
            show: false,
            create: 'create',
            edit: 'edit'
        },
        panelHeaderCreate: 'Создать',
        panelHeaderEdit: 'Изменить',
        panelSubmitButtonTextCreate: 'Создать',
        panelSubmitButtonTextEdit: 'Изменить',
        currentSetting: {
            item: new {Entity}(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default {entity}Data;

`