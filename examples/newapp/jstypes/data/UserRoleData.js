
	import {UserRole, UserRoleFilter} from '../apiModel';

let userRoleData = function () {
    return {
        userRoleFilter: new UserRoleFilter(),
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
        currentUserRoleItem: {
            item: new UserRole(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default userRoleData;

