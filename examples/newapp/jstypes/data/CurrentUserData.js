
	import {CurrentUser, CurrentUserFilter} from '../apiModel';

let currentUserData = function () {
    return {
        currentUserFilter: new CurrentUserFilter(),
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
        currentCurrentUserItem: {
            item: new CurrentUser(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default currentUserData;

