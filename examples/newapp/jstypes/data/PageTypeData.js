
	import {PageType, PageTypeFilter} from '../apiModel';

let pageTypeData = function () {
    return {
        pageTypeFilter: new PageTypeFilter(),
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
        currentPageTypeItem: {
            item: new PageType(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default pageTypeData;

