
	import {PageContentFilter, PageContentFilterFilter} from '../apiModel';

let pageContentFilterData = function () {
    return {
        pageContentFilterFilter: new PageContentFilterFilter(),
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
        currentPageContentFilterItem: {
            item: new PageContentFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default pageContentFilterData;

