using Pandora.Definitions.Attributes;
using Pandora.Definitions.CustomTypes;
using Pandora.Definitions.Interfaces;
using Pandora.Definitions.Operations;
using System;
using System.Collections.Generic;
using System.Net;

namespace Pandora.Definitions.ResourceManager.EventGrid.v2020_10_15_preview.EventSubscriptions
{
    internal class PartnerTopicEventSubscriptionsGetOperation : Operations.GetOperation
    {
        public override ResourceID? ResourceId() => new PartnerTopicEventSubscriptionId();

        public override Type? ResponseObject() => typeof(EventSubscriptionModel);


    }
}
