
	import {UserRoleFilter, UserRoleFilterFilter} from '../apiModel';

let userRoleFilterData = function () {
    return {
        userRoleFilterFilter: new UserRoleFilterFilter(),
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
        currentUserRoleFilterItem: {
            item: new UserRoleFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default userRoleFilterData;

