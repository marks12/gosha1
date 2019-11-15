
	import {UserFilter, UserFilterFilter} from '../apiModel';

let userFilterData = function () {
    return {
        userFilterFilter: new UserFilterFilter(),
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
        currentUserFilterItem: {
            item: new UserFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default userFilterData;

