
	import {PageTypeFilter, PageTypeFilterFilter} from '../apiModel';

let pageTypeFilterData = function () {
    return {
        pageTypeFilterFilter: new PageTypeFilterFilter(),
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
        currentPageTypeFilterItem: {
            item: new PageTypeFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default pageTypeFilterData;

