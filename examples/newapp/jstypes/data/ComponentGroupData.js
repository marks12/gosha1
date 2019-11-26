
	import {ComponentGroup, ComponentGroupFilter} from '../apiModel';

let componentGroupData = function () {
    return {
        componentGroupFilter: new ComponentGroupFilter(),
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
        currentComponentGroupItem: {
            item: new ComponentGroup(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default componentGroupData;

