
	import {RoleResourceFilter, RoleResourceFilterFilter} from '../apiModel';

let roleResourceFilterData = function () {
    return {
        roleResourceFilterFilter: new RoleResourceFilterFilter(),
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
        currentRoleResourceFilterItem: {
            item: new RoleResourceFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default roleResourceFilterData;

