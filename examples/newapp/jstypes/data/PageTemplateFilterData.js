
	import {PageTemplateFilter, PageTemplateFilterFilter} from '../apiModel';

let pageTemplateFilterData = function () {
    return {
        pageTemplateFilterFilter: new PageTemplateFilterFilter(),
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
        currentPageTemplateFilterItem: {
            item: new PageTemplateFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default pageTemplateFilterData;

