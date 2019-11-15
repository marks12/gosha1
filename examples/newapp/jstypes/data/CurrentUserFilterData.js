
	import {CurrentUserFilter, CurrentUserFilterFilter} from '../apiModel';

let currentUserFilterData = function () {
    return {
        currentUserFilterFilter: new CurrentUserFilterFilter(),
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
        currentCurrentUserFilterItem: {
            item: new CurrentUserFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default currentUserFilterData;

