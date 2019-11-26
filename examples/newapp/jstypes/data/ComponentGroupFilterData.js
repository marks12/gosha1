
	import {ComponentGroupFilter, ComponentGroupFilterFilter} from '../apiModel';

let componentGroupFilterData = function () {
    return {
        componentGroupFilterFilter: new ComponentGroupFilterFilter(),
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
        currentComponentGroupFilterItem: {
            item: new ComponentGroupFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default componentGroupFilterData;

