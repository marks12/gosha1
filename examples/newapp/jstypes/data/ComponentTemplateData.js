
	import {ComponentTemplate, ComponentTemplateFilter} from '../apiModel';

let componentTemplateData = function () {
    return {
        componentTemplateFilter: new ComponentTemplateFilter(),
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
        currentComponentTemplateItem: {
            item: new ComponentTemplate(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default componentTemplateData;

