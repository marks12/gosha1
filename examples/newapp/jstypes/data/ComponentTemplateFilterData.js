
	import {ComponentTemplateFilter, ComponentTemplateFilterFilter} from '../apiModel';

let componentTemplateFilterData = function () {
    return {
        componentTemplateFilterFilter: new ComponentTemplateFilterFilter(),
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
        currentComponentTemplateFilterItem: {
            item: new ComponentTemplateFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default componentTemplateFilterData;

