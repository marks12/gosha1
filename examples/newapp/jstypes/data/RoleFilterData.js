
	import {RoleFilter, RoleFilterFilter} from '../apiModel';

let roleFilterData = function () {
    return {
        roleFilterFilter: new RoleFilterFilter(),
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
        currentRoleFilterItem: {
            item: new RoleFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default roleFilterData;

