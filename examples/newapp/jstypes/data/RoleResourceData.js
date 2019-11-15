
	import {RoleResource, RoleResourceFilter} from '../apiModel';

let roleResourceData = function () {
    return {
        roleResourceFilter: new RoleResourceFilter(),
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
        currentRoleResourceItem: {
            item: new RoleResource(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default roleResourceData;

